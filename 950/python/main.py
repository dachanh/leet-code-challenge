from typing import List
from collections import deque

class Solution:
    def deckRevealedIncreasing(self, deck: List[int]) -> List[int]:
        N = len(deck)
        sorted_deck = sorted(deck)
        index = deque(range(N))
        result = [None] * N

        for card in sorted_deck:
            result[index.popleft()] = card
            if index:
                index.append(index.popleft())

        return result