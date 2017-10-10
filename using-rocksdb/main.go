package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/tecbot/gorocksdb"
)

func main() {
	// Open the RocksDB database. Create it new if it doesn't already exist.
	opts := gorocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(true)

	db, err := gorocksdb.OpenDb(opts, "using-rocksdb-sample-db")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
	// End of opening the RocksDB database.

	// Write to the RocksDB database.
	wo := gorocksdb.NewDefaultWriteOptions()
	err = db.Put(wo, []byte("message:1"), []byte("this is a message I've stored in Rocks DB!"))

	if err != nil {
		log.Fatal(err)
	}
	// End of writing to the RocksDB database.

	// Read a specific key from the RocksDB database.
	fmt.Println("beginning reading from the RocksDB database")
	ro := gorocksdb.NewDefaultReadOptions()
	output, err := db.Get(ro, []byte("message:1"))

	if err != nil {
		log.Fatal(err)
	}

	storedMessage := string(output.Data())
	output.Free()

	fmt.Println(storedMessage)
	fmt.Println("finished doing a basic read from the RocksDB database")
	fmt.Println()
	// End of reading from the RocksDB database.

	// Write a whole bunch of structured keys to the RocksDB database for seeking.
	err = db.Put(wo, []byte("JOSEPH CAMPBELL:1"), []byte("Well, because there is a certain typical hero sequence of actions, which can be detected in stories from all over the world, and from many, many periods of history. And I think it’s essentially, you might say, the one deed done by many, many different people."))

	if err != nil {
		log.Fatal(err)
	}

	err = db.Put(wo, []byte("BILL MOYERS:1"), []byte("Why are there so many stories of the hero or of heroes in mythology?"))

	if err != nil {
		log.Fatal(err)
	}

	err = db.Put(wo, []byte("JOSEPH CAMPBELL:2"), []byte("Well, because that’s what’s worth writing about. I mean, even in popular novel writing, you see, these the main character is the hero or heroine, that is to say, someone who has found or achieved or done something beyond the normal range of achievement and experience. A hero properly is someone who has given his life to something bigger than himself or other than himself."))

	if err != nil {
		log.Fatal(err)
	}

	err = db.Put(wo, []byte("BILL MOYERS:2"), []byte("So in all of these cultures, whatever the costume the hero might be wearing, what is the deed?"))

	if err != nil {
		log.Fatal(err)
	}

	err = db.Put(wo, []byte("JOSEPH CAMPBELL:3"), []byte("Well, there are two types of deed. One is the physical deed; the hero who has performed a war act or a physical act of heroism ñ saving a life, that’s a hero act. Giving himself, sacrificing himself to another. And the other kind is the spiritual hero, who has learned or found a mode of experiencing the supernormal range of human spiritual life, and then come back and communicated it. It’s a cycle, it’s a going and a return, that the hero cycle represents."))

	if err != nil {
		log.Fatal(err)
	}

	err = db.Put(wo, []byte("JOSEPH CAMPBELL:4"), []byte("But then this can be seen also in the simple initiation ritual, where a child has to give up his childhood and become an adult, has to die, you might say, to his infantile personality and psyche and come back as a self-responsible adult. It’s a fundamental experience that everyone has to undergo, where in our childhood for at least 14 years, and then to get out of that posture of dependency, psychological dependency, into one of psychological self-responsibility, requires a death and resurrection, and that is the basic motif of the hero journey, Leaving one condition, finding the source of life to bring you forth in a richer or more mature or other condition."))

	if err != nil {
		log.Fatal(err)
	}

	err = db.Put(wo, []byte("BILL MOYERS:3"), []byte("So that if we happen not to be heroes in the grand sense of redeeming society, we have to lake that journey ourselves, spiritually, psychologically, inside us."))

	if err != nil {
		log.Fatal(err)
	}
	// End of writing a bunch of keys to the RocksDB database.

	// Begin iterating through keys but specify a structured key prefix to start at.
	iter := db.NewIterator(ro)

	fmt.Println("begin iteration")
	for iter.Seek([]byte("JOSEPH CAMPBELL:")); iter.Valid(); iter.Next() {
		keyData := iter.Key()
		valueData := iter.Value()

		key := string(keyData.Data())
		value := string(valueData.Data())

		keyData.Free()
		valueData.Free()

		if !strings.HasPrefix(key, "JOSEPH CAMPBELL:") {
			// If the prefix doesn't equal "JOSEPH CAMPBELL:" we'll end immediately.
			break
		}

		fmt.Println("key:", key)
		fmt.Println("value:", value)
		fmt.Println()
	}
	fmt.Println("ending iteration")
	// End of iterating through keys.
}
