class MyQueue:

    def __init__(self):
        self.data = []
        self.l = 0

    def push(self, x: int) -> None:
        self.data.append(x)
        self.l += 1

    def pop(self) -> int:
        if self.empty():
            return None

        res = self.data[0]
        self.data = self.data[1:]
        self.l -= 1
        return res

    def peek(self) -> int:
      if self.empty():
        return None
      return self.data[0]

    def empty(self) -> bool:
        return self.l == 0