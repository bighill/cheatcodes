# python

Serve website locally

```bash
# python 2.x
python -m SimpleHTTPServer

# python 3.x
python3 -m http.server
```

## Lists

```python
integer_list = [1, 2, 3]
heterogeneous_list = ["string", 0.1, True]
list_of_lists = [integer_list, heterogeneous_list, []]
list_length = len(integer_list)
list_sum = sum(integer_list)
```

Get from lists

```python
x = range(10)
zero = x[0]   # first in list
one = x[1]
nine = x[-1]  # last in list
eight = x[-2]
```

Slice list

```python
x = range(10)
first_three = x[:3]
three_to_end = x[3:]
one_to_four = x[1:5]
last_three = x[-3:]
without_first_and_last = x[1:-1]
copy_of_x = x[:]
```

Find in list

```python
1 in [1, 2, 3]  # True
0 in [1, 2, 3]  # False
```

Sorting

```python
x = [4, 1, 2, 3]
y = sorted(x) # sorted() will NOT mutate x
x.sort()      # .sort() WILL mutate x
```

Transform list

```python
even_numbers = [x for x in range(5) if x % 2 == 0]
squares = [x * x for x in range(5)]
even_squares = [x * x for x in even_numbers]
```

Enumerate when we want the element & index

```python
x = ['red', 'blue', 'green']

for i, color in enumerate(x):
    print(i, color)
```

Zip

```python
list1 = ['a', 'b', 'c']
list2 = [1, 2, 3]
pairs = zip(list1, list2)     # is [('a', 1), ('b', 2), ('c', 3)]
letters, number = zip(*pairs) # this will UNzip
```

## Tuples

```python
def sum_and_product(x, y):
    """Tuples are a convenient way to return multiple values from functions"""
    return (x + y), (x * y)

sp = sum_and_product(2, 3)  # equals (5, 6)
s, p = sum_and_product(5, 10)  # s is 15, p is 50
```

## Dictionaries

```python
grades = {"Joel": 80, "Tim": 95}
```

Get from dictionary

```python
joels_grade = grades.get("Joel", 0)
kates_grade = grades.get("Kate", 0)   # 0 is the default
no_ones_grade = grades.get("No One")  # default default is None
```

Keys, values, tuples

```python
grade_keys = grades.keys()      # list of keys
grade_values = grades.values()  # list of values
grade_items = grades.items()    # list of tuples
```

Transform list to dict

```python
square_dict = {x: x * x for x in range(5)}
```

## Sets

```python
s = set()
s.add(1)
foo_list = [1, 2, 3]
foo_set = set(foo_list)
```

Transform list to set

```python
square_set = {x * x for x in [1, -1]}
```

## Map

```python
def double(x): return 2*x

xs = [1, 2, 3, 4]

# map - these two lines are equivelent
twice_xs = [double(x) for x in xs]
twice_xs = map(double, xs)
```

# Filter

```python
def is_even(x): return x % 2 == 0

xs = [1, 2, 3, 4]

# filter - these two lines are equivelent
x_evens = [x for x in xs if is_even(x)]
x_evens = filter(is_even, xs)
```

# args & kwargs

```python
def magic(*args, **kwargs):
    print(f"unnamed args: {args}")
    print(f"keyword args: {kwargs}")


magic(1, 2, key="word", key2="word2")
# unnamed args: (1, 2)
# keyword args: {'key2': 'word2', 'key': 'word'}
```
