"""
This module contains a simple version of TF-IDF text processor.
"""
from math import log
from typing import TypeVar
from numpy import array

Func = TypeVar("Func")


def tf_func(term, doc: str) -> float:
    """
    Find the importance rate of the term
    in the specified document.
    """
    return doc.lower().count(term) / len(doc.split())


def idf_func(term: str, d_corpus: list[str]) -> float:
    """
    Find the importance of the term across
    the document corpus.
    """
    return log(len(d_corpus) + 1 / len([d for d in d_corpus if d.lower().count(term) > 0]) + 1) + 1


def word_importance(words, d_corpus: list[str], tf: Func, idf: Func) -> dict[str, float]:
    scores: dict[str, float] = {}
    for i in range(len(words)):
        scores[words[i]] = tf(words[i], d_corpus[2]) * idf(words[i], d_corpus)

    return scores

def vectorize_scores(scores: dict[str, float]) -> list[float]:
	return array([v for v in scores.values()])

if __name__ == '__main__':
    doc = "this sentence is about dog, the dog is a very friendly animal"
    doc1 = "this sentence in not about dog, the cat is lazier and faster than dog, since a cat can communicate to another cat"
    doc2 = "this sentence is neither about cat nor dog, it's about food, the food in very essential to animals"
    doc3 = "this sentence is not about food, it's about how to feed animals by their food"
    doc4 = "the previous sentence has more words about dog."
    doc5 = "the first sentence is about dog but this one is about cat, because cat is a domestic animal as well."
    corpus = [doc, doc1, doc2, doc3, doc4, doc5]
    categs = ["cat", "dog", "animal", "food", "sentence"]
    
    scores = word_importance(categs, corpus, tf_func, idf_func)
    
    print(vectorize_scores(scores))
	
