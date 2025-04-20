# given an array of integers, integers represent heigts
# {4,2,3,1}

# return an array, of integers, representign arrays that are taller than all the elements to their rights
# the answer for the above would be {0,2,3}

# heights = {4,3,2,1}
# result = {0,1,2,3}




# 1 2 3 4

def seeable_indexes(heights: list[int]) -> list[int]:
    heights = [4,3,2,1]
    res = []

    curr_max = float('-inf') # 4

    # 4, 3, 2, 1
    # i = -1
    i = len(heights) - 1

    while(i >= 0):
        if heights[i] > curr_max:
            curr_max = heights[i]
            res.append(i)
        i -= 1
    
    return res[::-1]

# Time O(n)
# Space O(n)
heights = [4, 3, 2, 1]
print(seeable_indexes(heights))
