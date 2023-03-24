# 将使⽤ pytest 执⾏测试
from main import solution


def test_solution_1():
    expected = 1
    nums = [1, 2, 3]
    target = 2
    assert solution(nums, target) == expected


def test_solution_2():
    expected = 0
    nums = [1, 3, 5, 7, 9]
    target = 10
    assert solution(nums, target) == expected


def test_solution_3():
    expected = -1
    nums = [2, 4, 6, 8, 10]
    target = 1
    assert solution(nums, target) == expected
