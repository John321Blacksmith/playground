"""
This module contains a simple version of TF-IDF text processor.
"""
from math import log
from dataclasses import dataclass
from typing import TypeVar, List, Dict
from numpy import array


Func = TypeVar("Func")


text = "this sentence is about dog, the dog is a very friendly animal. this sentence in not about dog, the cat is lazier and faster than dog, since a cat can communicate to another cat. this sentence is neither about cat nor dog, it's about food, the food in very essential to animals. this sentence is not about food, it's about how to feed animals by their food. the previous sentence has more words about dog. the first sentence is about dog but this one is about cat, because cat is a domestic animal as well."


categs = ["cat", "dog", "communic", "animal", "food", "to", "sentence", "is", "first", "shit"]


indexes: dict[float, set[str]] = {
	1.0: {"dog", "bark", "bit", "doggy", "fluffy", "pet", "wolf", "wolves"},
	2.0: {"cat", "meow", "kitty", "fluffy"},
	3.0: {"food", "canteen", "kitchen", "delicious", "sweet", "salt", "spicy"},
}


categ_maps: dict[str, float] = {
	"dog": indexes[1.0],
	"cat": indexes[2.0],
	"food": indexes[3.0],
}


@dataclass
class Sentence:
	category: float
	tokens: set[str]
	
	def __len__(self) -> int:
		return len(self.tokens)
		
	def count(self, term: str) -> int:
		return len([t for t in self.tokens if t.startswith(term)])
	

def split_to_sentences(inp: str) -> List[Sentence]:
	return [
		Sentence(category=.0, tokens={w for w in s.split(" ") if len(w) > 2})
				for s in inp.split(".") if len(s) > 0
			]
				

def tf(term: str, sentence: Sentence) -> float:
	return sentence.count(term) / len(sentence)
	

def idf(term: str, sentences: List[Sentence]) -> float:
	return log((len(sentences) + 1) / (len([s for s in sentences if s.count(term) > 0]) + 1)) + 1


def get_input_scores_map(categs: List[str], sentences: List[Sentence], tf: Func, idf: Func) -> Dict[str, float]:
	scores: Dict[str, float] = {}
	for i in range(len(categs)):
		for j in range(len(sentences)):
			if categs[i] not in scores:
				scores[categs[i]] = 0
			scores[categs[i]] += tf(categs[i], sentences[j]) * idf(categs[i], sentences)
	return scores


def get_sentence_scores_map(
		categs: List[str],
		sentence: Sentence,
		sentences: List[Sentence],
		tf: Func,
		idf: Func
	) -> Dict[str, float]:
	scores: Dict[str, float] = {}
	for i in range(len(categs)):
		if categs[i] not in scores:
			scores[categs[i]] = 0
		scores[categs[i]] = tf(categs[i], sentence) * idf(categs[i], sentences)
		
	return scores


def vectorize_scores(sentence_scores: Dict[str, float]) -> List[float]:
	return array([v for v in sentence_scores.values()])


def aggregate_vectors(categs: List[str], sentences: List[Sentence]) -> List[List[float]]:
	...
	

def main():
	vectors: List[List[float]] = []
	sum_vectors = array([.0 for i in range(len(categs))])
	sentences = split_to_sentences(text)
	
	for i in range(len(sentences)):
		scores = get_sentence_scores_map(categs, sentences[i], sentences, tf, idf)
		vector = vectorize_scores(scores)
		vectors.append(vector)

	for i in range(len(vectors)):
		sum_vectors += vectors[i]
	

def sort(pairs: List[tuple[str, float]]):
	if len(pairs) < 2:
		return pairs
		
	else:
		pivot = pairs[len(pairs) // 2]
		smaller = [p for p in pairs if p[1] < pivot[1]]
		greater = [p for p in pairs if p[1] > pivot[1]]
		
		return sort(greater) + [pivot] + sort(smaller)
	

if __name__ == '__main__':
	tokens_bag: list[list[str]] = [[t for t in s.split(" ") if len(t) > 2] for s in text.split(".") if len(s) > 0]
	sentences = [Sentence(category="", tokens=s_token) for s_token in tokens_bag]
	vectors: list[list[str]] = []

	for i in range(len(sentences)):
		vectors.append([v for v in get_sentence_scores_map(categs, sentences[i], sentences, tf, idf).values()])

	print(vectors)
	
