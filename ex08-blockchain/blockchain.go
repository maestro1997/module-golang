package main

import (
	"time"
	"fmt"
	"bytes"
	"strconv"
	"crypto/sha256"
	"os"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func AppendBlock(db *sql.DB, b *Block) {
	db.Exec("INSERT INTO blockchain (timestamp,data,hash,prevHash)	values ($1,$2,$3,$4)",
		b.Timestamp, b.Data, b.Hash, b.PrevBlockHash)
}

func CreateTable(db *sql.DB) {
	db.Exec("CREATE TABLE IF NOT EXISTS blockchain (id INTEGER PRIMARY KEY AUTOINCREMENT, timestamp INTEGER, data TEXT, hash TEXT,prevHash TEXT)")
	row :=db.QueryRow("SELECT COUNT(*) FROM blockchain")
	count := 0
	row.Scan(&count)
	if count == 0 {
		AppendBlock(db, NewGenesisBlock())
	}
}

func getPrevHash(db *sql.DB) []byte {
	row := db.QueryRow("SELECT hash FROM blockchain WHERE id=(SELECT MAX(id) FROM blockchain)")
	var prevHash []byte
	row.Scan(&prevHash)
	return prevHash
}

func AddBlock(db *sql.DB, data string) {
	prevHash := getPrevHash(db)
	b := NewBlock(data, prevHash)
	b.PrevBlockHash = prevHash
	AppendBlock(db, b)
}

func PrintBlockChain(db *sql.DB) {
	rows, _ := db.Query("SELECT id,timestamp,data FROM blockchain")
	defer rows.Close()
	var timestamp int64
	var data []byte
	var id int;
	for rows.Next() {
		rows.Scan(&id,&timestamp,&data)
		fmt.Printf("Block #%d : timestamp = %d, data = %p\n",id,timestamp,data)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You didn,t choose action!")
		return
	}
	var data string
	db, err := sql.Open("sqlite3", "bc.db")
	if err != nil {
		fmt.Println("DB opening error")
		return
	}
	defer db.Close()
	CreateTable(db)
	if os.Args[1] == "add" {
		fmt.Print("Please, take data : ")
		fmt.Scanf("%s", &data)
		AddBlock(db, data)
		return
	}
	if os.Args[1] == "list" {
		PrintBlockChain(db)
		return
	}
	fmt.Println("Error: unknown command")
}


