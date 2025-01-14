package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	miner = "0xd912aecb07e9f4e1ea8e6b4779e7fb6aa1c3e4d8"
	// SPDX-License-Identifier: UNLICENSED
	// pragma solidity ^0.7.0;
	// contract Bribe {
	//     function bribe() payable public {
	//         block.coinbase.transfer(msg.value);
	//     }
	// }
	bribeContractBin = `0x6080604052348015600f57600080fd5b5060a78061001e6000396000f3fe608060405260043610601c5760003560e01c806337d0208c146021575b600080fd5b60276029565b005b4173ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f19350505050158015606e573d6000803e3d6000fd5b5056fea2646970667358221220862610b9326c9523da6465ba88229d2c6b26ff844b3c9cddb807c2c1ab401dd964736f6c63430007050033`
	bribeContractABI = `[
    {
      inputs: [],
      name: 'bribe',
      outputs: [],
      stateMutability: 'payable',
      type: 'function'
    }
  ]
`
)

var (
	clientDial = flag.String(
		"client_dial", "ws://127.0.0.1:8546", "could be websocket or IPC",
	)
	at        = flag.Uint64("kickoff", 2, "what number to kick off at")
	faucet, _ = crypto.HexToECDSA(
		"133be114715e5fe528a1b8adf36792160601a2d63ab59d1fd454275b31328791",
	)
	keys        = []*ecdsa.PrivateKey{faucet}
	bribeABI, _ = abi.JSON(strings.NewReader(string(bribeContractABI)))
)

func banditAttack(
	client *ethclient.Client,
	toAddr common.Address,
	chainID *big.Int,
) ([]types.Transactions, error) {

	packed, err := bribeABI.Methods["bribe"].Inputs.Pack()
	if err != nil {
		return nil, err
	}
	txs := make(types.Transactions, len(keys))

	for i, key := range keys {
		k := crypto.PubkeyToAddress(key.PublicKey)
		non, err := client.NonceAt(
			context.Background(), k, nil,
		)
		if err != nil {
			return nil, err
		}

		balance, err := client.BalanceAt(context.Background(), k, nil)
		if err != nil {
			return nil, err
		}
		if balance.Cmp(common.Big0) == 0 {
			return nil, errors.New("need non-zero balance")
		}
		t := types.NewTransaction(
			non,
			toAddr,
			new(big.Int),
			100_000,
			big.NewInt(3e9),
			packed,
		)
		t, err = types.SignTx(t, types.NewEIP155Signer(chainID), key)
		if err != nil {
			return nil, err
		}
		txs[i] = t
	}
	return []types.Transactions{txs, txs, txs}, nil
}

func deployBribeContract(
	client *ethclient.Client,
	chainID *big.Int,
) (*types.Transaction, error) {
	t := types.NewContractCreation(
		0, new(big.Int), 400_000, big.NewInt(10e9),
		common.Hex2Bytes(bribeContractBin),
	)

	t, err := types.SignTx(t, types.NewEIP155Signer(chainID), faucet)
	if err != nil {
		return nil, err
	}

	return t, client.SendTransaction(context.Background(), t)
}

func program() error {
	client, err := ethclient.Dial(*clientDial)
	if err != nil {
		return err
	}

	ch := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(
		context.Background(), ch,
	)

	if err != nil {
		return err
	}

	var newContractAddr common.Address

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return err
	}
	deployAt := *at

	for {
		select {
		case e := <-sub.Err():
			return e
		case incoming := <-ch:
			blockNumber := incoming.Number.Uint64()
			if blockNumber == deployAt {
				t, err := deployBribeContract(client, chainID)
				if err != nil {
					return err
				}

				newContractAddr = crypto.CreateAddress(
					crypto.PubkeyToAddress(faucet.PublicKey),
					t.Nonce(),
				)
				fmt.Println("\tdeployed bribe contract ", newContractAddr.Hex(), blockNumber)
				continue
			}

			fmt.Println(
				"new head", blockNumber, incoming.Hash(),
			)

			if blockNumber == deployAt+1 {
				usedTxs, err := banditAttack(client, newContractAddr, chainID)
				if err != nil {
					return err
				}

				fmt.Println("using as parent hash",
					incoming.Hash().Hex(), incoming.Number,
				)

				init := time.Now()
				empty := common.Hash{}

				if err := client.SendBandit(
					context.Background(), []*types.Bandit{
						&types.Bandit{
							Transactions: usedTxs[0],
							Timestamp:    uint64(init.Add(-time.Second * 45).Unix()),
							CoinbaseDiff: 1e17,
							Number:       blockNumber - 2,
							ParentHash:   empty,
						},
						&types.Bandit{
							Transactions: usedTxs[1],
							Timestamp:    uint64(init.Add(-time.Second * 30).Unix()),
							CoinbaseDiff: 1e17,
							Number:       blockNumber - 1,
							ParentHash:   empty,
						},
						&types.Bandit{
							Transactions: usedTxs[2],
							Timestamp:    uint64(init.Add(-time.Second * 15).Unix()),
							CoinbaseDiff: 1e17,
							Number:       blockNumber - 0,
							ParentHash:   empty,
						},
					},
				); err != nil {
					return err
				}
				fmt.Println("kicked off time bandit")
			}
		}
	}
}

func main() {
	flag.Parse()
	if err := program(); err != nil {
		log.Fatal(err)
	}
}
