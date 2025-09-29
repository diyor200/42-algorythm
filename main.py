r = [1,2,3,4,5]
print(r[:len(r)-2])
class MyStack:

    def __init__(self):
        self.data = list()
        self.l = 0
        

    def push(self, x: int) -> None:
        self.data.append(x)
        self.l += 1
        return

    def pop(self) -> int:
        if self.l == 0:
            return None
        
        res = self.data[self.l-1]
        self.data = self.data[:self.l-2]
        self.l -= 1

        return res

    def top(self) -> int:
        if self.l == 0:
            return None
        
        return self.data[self.l-1]
        

    def empty(self) -> bool:
        return self.l == 0
