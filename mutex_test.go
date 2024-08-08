package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	var x int
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Hasilnya adalah", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (t *BankAccount) AddBalance(amount int) {
	t.RWMutex.Lock()
	t.Balance = t.Balance + amount
	t.RWMutex.Unlock()
}

func (t *BankAccount) GetBalance() int {
	//t.RWMutex.RLock()
	balance := t.Balance
	//t.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Hasilnya adalah", account.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance += amount
}

func Transfer(from, to *UserBalance, amount int) {
	from.Lock()
	fmt.Println("lock user:", from.Name)
	from.Change(-amount)
	time.Sleep(1 * time.Second)

	to.Lock()
	fmt.Println("lock user:", to.Name)
	to.Change(amount)
	time.Sleep(1 * time.Second)

	from.Unlock()
	to.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Deni",
		Balance: 1_000_000,
	}

	user2 := UserBalance{
		Name:    "Budi",
		Balance: 1_000_000,
	}

	go Transfer(&user1, &user2, 100_000)
	go Transfer(&user2, &user1, 200_000)

	time.Sleep(10 * time.Second)
	//deadlock occurred, each user waiting for other user to unlock

	fmt.Println("User:", user1.Name, "Balance:", user1.Balance)
	fmt.Println("User:", user2.Name, "Balance:", user2.Balance)
}
