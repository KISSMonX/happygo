// this is a file that puts together all redigo examples for convenience
// (see https://godoc.org/github.com/gomodule/redigo/redis#pkg-examples)
//
// start by ensuring that redis is running on port 6379 (`redis-server`)
// uncomment the main method as needed, and run the script (`go run main.go`)
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	c     redis.Conn
	err   error
	reply interface{}
)

func init() {
	c, err = redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer c.Close()
	c.Do("FLUSHALL")
	argsExample()
	//boolExample()
	//intExample()
	//intsExample()
	//scanExample()
	//scanSliceExample()
	//stringExample()
}

type Timestamp time.Time

func (t *Timestamp) RedisScan(x interface{}) error {
	bs, ok := x.([]byte)
	if !ok {
		return fmt.Errorf("expected []byte, got %T", x)
	}
	tt, err := time.Parse(time.RFC3339, string(bs))
	if err != nil {
		return err
	}
	*t = Timestamp(tt)
	return nil
}

func argsExample() {

	var p1, p2 struct {
		Title  string    `redis:"title"`
		Author string    `redis:"author"`
		Body   string    `redis:"body"`
		Age    int       `redis:"age"`
		Hight  float64   `redis:"hight"`
		Time   time.Time `redis:"time"`
	}

	p1.Title = "Example"
	p1.Author = "Gary"
	p1.Body = "Hello"
	p1.Age = 233
	p1.Hight = 242.325365
	p1.Time = time.Now()

	if _, err := c.Do("HMSET", redis.Args{}.Add("id1").AddFlat(&p1)...); err != nil {
		fmt.Println(err)
		return
	}

	m := map[string]interface{}{
		"title":  "Example2",
		"author": "Steve",
		"body":   "Map",
		"age":    123245325,
		"hight":  24135.6457547547547,
		"time":   time.Now(),
	}

	if _, err := c.Do("HMSET", redis.Args{}.Add("id2").AddFlat(m)...); err != nil {
		fmt.Println(err)
		return
	}

	for _, id := range []string{"id1", "id2"} {

		v, err := redis.Values(c.Do("HGETALL", id))
		if err != nil {
			fmt.Println(err)
			return
		}

		if err := redis.ScanStruct(v, &p2); err != nil {
			fmt.Println(err)
			return
		}

		b, err := json.Marshal(p2)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b))
	}
}

func boolExample() {
	c.Do("SET", "foo", 1)
	exists, _ := redis.Bool(c.Do("EXISTS", "foo"))
	fmt.Printf("%#v\n", exists)
}

func intExample() {
	c.Do("SET", "k1", 1)
	n, _ := redis.Int(c.Do("GET", "k1"))
	fmt.Printf("%#v\n", n)
	n, _ = redis.Int(c.Do("INCR", "k1"))
	fmt.Printf("%#v\n", n)
}

func intsExample() {
	c.Do("SADD", "set_with_integers", 4, 5, 6)
	ints, _ := redis.Ints(c.Do("SMEMBERS", "set_with_integers"))
	fmt.Printf("%#v\n", ints)
}

func scanExample() {
	c.Send("HMSET", "album:1", "title", "Red", "rating", 5)
	c.Send("HMSET", "album:2", "title", "Earthbound", "rating", 1)
	c.Send("HMSET", "album:3", "title", "Beat")
	c.Send("LPUSH", "albums", "1")
	c.Send("LPUSH", "albums", "2")
	c.Send("LPUSH", "albums", "3")
	values, err := redis.Values(c.Do("SORT", "albums",
		"BY", "album:*->rating",
		"GET", "album:*->title",
		"GET", "album:*->rating"))
	if err != nil {
		fmt.Println(err)
		return
	}

	for len(values) > 0 {
		var title string
		rating := -1 // initialize to illegal value to detect nil.
		values, err = redis.Scan(values, &title, &rating)
		if err != nil {
			fmt.Println(err)
			return
		}
		if rating == -1 {
			fmt.Println(title, "not-rated")
		} else {
			fmt.Println(title, rating)
		}
	}
}

func scanSliceExample() {
	c.Send("HMSET", "album:1", "title", "Red", "rating", 5)
	c.Send("HMSET", "album:2", "title", "Earthbound", "rating", 1)
	c.Send("HMSET", "album:3", "title", "Beat", "rating", 4)
	c.Send("LPUSH", "albums", "1")
	c.Send("LPUSH", "albums", "2")
	c.Send("LPUSH", "albums", "3")
	values, err := redis.Values(c.Do("SORT", "albums",
		"BY", "album:*->rating",
		"GET", "album:*->title",
		"GET", "album:*->rating"))
	if err != nil {
		fmt.Println(err)
		return
	}

	var albums []struct {
		Title  string
		Rating int
	}
	if err := redis.ScanSlice(values, &albums); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", albums)
}

func stringExample() {
	c.Do("SET", "hello", "world")
	s, err := redis.String(c.Do("GET", "hello"))
	fmt.Printf("%#v %v\n", s, err)
}
