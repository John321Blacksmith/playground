package main

import (
	"fmt"
	"math"
	"strings"
)

func tf_func(term string, sentence string) float64 {
	return float64(strings.Count(term, sentence) / len(strings.Split(sentence, "")))
}

func idf_func(term string, d_corpus []string) float64 {
	sentences := []string{}
	for i := range len(d_corpus) {
		if strings.Count(d_corpus[i], term) > 0 {
			sentences = append(sentences, d_corpus[i])
		}
	}
	quotent := float64((len(d_corpus) + 1) / (len(sentences) + 1))
	return math.Log(quotent) + 1
}

func getrWordImportance(categs, d_corpus []string) map[string]float64 {
	var scores map[string]float64 = map[string]float64{}

	for i := range len(categs) {
		fmt.Println(tf_func(categs[i], d_corpus[0]))
		scores[categs[i]] = tf_func(categs[i], d_corpus[0]) * idf_func(categs[i], d_corpus)
	}
	return scores
}

func main() {
	doc := "this sentence is about dog, the dog is a very friendly animal"
	doc1 := "this sentence in not about dog, the cat is lazier and faster than dog, since a cat can communicate to another cat"
	doc2 := "this sentence is neither about cat nor dog, it's about food, the food in very essential to animals"
	doc3 := "this sentence is not about food, it's about how to feed animals by their food"
	doc4 := "the previous sentence has more words about dog."
	doc5 := "the first sentence is about dog but this one is about cat, because cat is a domestic animal as well."

	corpus := []string{doc, doc1, doc2, doc3, doc4, doc5}
	categs := []string{"cat", "dog", "animal", "food", "sentence"}

	fmt.Println(getrWordImportance(categs, corpus))
}
