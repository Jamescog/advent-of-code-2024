"""Advent of code Day-2
"""

class Solution:
    """Solution"""

    def __init__(self) -> None:
        """Read the input data"""
        try:
            with open('input.txt') as f:
                data = f.readlines()
        except Exception as e:
            print("Error Reading file: ", e)

        report = []
        for line in data:
            l = [int(i) for i in line.split(" ")]
            report.append(l)

        self.report = report

    def __is_gradual(self, r: list) -> bool:
        """Returns if the levels are either all increasing or all decreasing."""
        t = []
        for i in range(1, len(r)):
            t.append(r[i] > r[i - 1])
        return True if len(set(t)) == 1 else False


    def __is_valid_diff(self, r:list) -> bool:
        """Returns if any two adjacent levels differ by at least one and at most three."""
        for i in range(1, len(r)):
            if abs(r[i] - r[i-1]) not in [1,2,3]:
                return False        
        return True
    

    def count_safe_reports(self) -> int:
        """Return how many reports are safe without dampener"""

        safe_count =0 
        for i in self.report:
            if self.__is_gradual(i) and self.__is_valid_diff(i):
                safe_count += 1
        return safe_count
    

    def __is_safe_with_dampener(self, r: list) -> bool:
        """Returns if the report is safe with dampener"""
        for i in range(len(r)):
            candidate = r[:i] + r[i + 1:]
            if self.__is_gradual(candidate) and self.__is_valid_diff(candidate):
                return True

        return False 

    
    def count_safe_reports_with_dampener(self) -> int:
        """Returns how many reports are safe with dampener"""

        count_safe = 0
        for i in self.report:
            if self.__is_gradual(i) and self.__is_valid_diff(i):
                count_safe += 1
            else:
                if self.__is_safe_with_dampener(i):
                    count_safe += 1

        return count_safe


if __name__ == "__main__":
    s = Solution()
    print(s.count_safe_reports())
    print(s.count_safe_reports_with_dampener())