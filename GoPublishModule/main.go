package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"

	// "rsc.io/quote/v3" // this is just for external dependency
)

type results struct {
	Result []aQuote
}

type aQuote struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

// this is just for external dependency
func Hello() string {
	// return quote.HelloV3()
	return "Hello!"
}

func quoteOfTheDay(w http.ResponseWriter, req *http.Request) {
	var res results
	var toReturn string

	// // workaround for CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	json.Unmarshal([]byte(Quotes), &res)
	thequote, _ := json.Marshal(res.Result[rand.Intn(len(res.Result))])
	toReturn = string(thequote)
	fmt.Println(toReturn)
	fmt.Fprintf(w, toReturn)
}

func main() {
	fqdn := "0.0.0.0"
	port := "8090"

	if tmp := os.Getenv("FQDN"); tmp != "" {
		fqdn = tmp
	}

	if tmp := os.Getenv("PORT"); tmp != "" {
		port = tmp
	}

	http.HandleFunc("/quote", quoteOfTheDay)
	http.ListenAndServe(fqdn+":"+port, nil)
}

const Quotes = `{
	"result": [
	  {
		"quote": "A good head and a good heart are always a formidable combination.",
		"author": "Nelson Mandela"
	  },
	  {
		"quote": "Education is the most powerful weapon which you can use to change the world.",
		"author": "Nelson Mandela"
	  },
	  {
		"quote": "Education is the most powerful weapon which you can use to change the world.",
		"author": "Nelson Mandela"
	  },
	  {
		"quote": "For to be free is not merely to cast off one's chains, but to live in a way that respects and enhances the freedom of others.",
		"author": "Nelson Mandela"
	  },
	  {
		"quote": "I am not a saint, unless you think of a saint as a sinner who keeps on trying.",
		"author": "Nelson Mandela"
	  },
	  {
		"quote": "I learned that courage is not the absence of fear but the triumph over it.",
		"author": "Nelson Mandela"
	  },
	  {
		"quote": "I learned that courage was not the absence of fear, but the triumph over it. The brave man is not he who does not feel afraid, but he who conquers that fear.",
		"author": "Nelson Mandela"
	  },
	  {
		"quote": "If you talk to a man in a language he understands, that goes to his head. If you talk to him in his language, that goes to his heart.",
		"author": "Nelson Mandela"
	  },
	  {
		"quote": "It always seems impossible until it's done.",
		"author": "Nelson Mandela"
	  },
	  {
		"quote": "Like slavery and apartheid, poverty is not natural. It is man-made and it can be overcome and eradicated by the actions of human beings.",
		"author": "Nelson Mandela"
	  },
	  {
		"quote": "May your choices reflect your hopes, not your fears.",
		"author": "Nelson Mandela"
	  },
	  {
		"quote": "Only free men can negotiate; prisoners cannot enter into contracts. Your freedom and mine cannot be separated.",
		"author": "Nelson Mandela"
	  },
	  {
		"quote": "Sometimes it falls upon a generation to be great. You can be that great generation.",
		"author": "Nelson Mandela"
	  },
	  {
		"quote": "The greatest glory in living lies not in never falling, but in rising every time we fall.",
		"author": "Nelson Mandela"
	  },
	  {
		"quote": "The greatest glory in living lies not in never falling, but in rising every time we fall.",
		"author": "Nelson Mandela"
	  },
	  {
		"quote": "There is nothing like returing to a place that remains uncganged to find how you yourself have altered.",
		"author": "Nelson Mandela"
	  },
	  {
		"quote": "There is nothing like returning to a place that remains unchanged to find the ways in which you yourself have altered.",
		"author": "Nelson Mandela"
	  },
	  {
		"quote": "When a boat come upon a  winding river, slow down and the river becomes straight.",
		"author": "Nelson Mandela"
	  },
	  {
		"quote": "“And as we let our own light shine, we unconsciously give other people permission to do the same”",
		"author": "Nelson Mandela"
	  },
	  {
		"quote": "“As I have said, the first thing is to be honest with yourself. You can never have an impact on society if you have not changed yourself... Great peacemakers are all people of integrity, of honesty, but humility.”",
		"author": "Nelson Mandela"
	  },
	  {
		"quote": "“If you want to make peace with your enemy, you have to work with your enemy. Then he becomes your partner.”",
		"author": "Nelson Mandela"
	  },
	  {
		"quote": "“It is better to lead from behind and to put others in front, especially when you celebrate victory when nice things occur. You take the front line when there is danger. Then people will appreciate your leadership.”",
		"author": "Nelson Mandela"
	  },
	  {
		"quote": "“The greatest glory in living lies not in never falling, but in rising every time we fall.”",
		"author": "Nelson Mandela"
	  },
	  {
		"quote": "“There is no passion to be found playing small - in settling for a life that is less than the one you are capable of living.”",
		"author": "Nelson Mandela"
	  }
	]
  }`
