"""
This module contains a simple version of TF-IDF text processor.
"""
from math import log
from typing import TypeVar

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
        scores[words[i]] = tf(words[i], d_corpus[3]) * idf(words[i], d_corpus)

    return scores


if __name__ == '__main__':
    doc0 = "this sentence is about dog, the dog is a very friendly animal"
    doc1 = "this sentence in not about dog, the cat is lazier and faster than dog, since a cat can communicate to another cat"
    doc2 = "this sentence is neither about cat nor dog, it's about food, the food in very essential to animals"
    doc3 = "this sentence is not about food, it's about how to feed animals by their food"

    corpus = [doc0, doc1, doc2, doc3]
    words = ["cat", "dog", "food"]
    print(word_importance(words, corpus, tf_func, idf_func))

