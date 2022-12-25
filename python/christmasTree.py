def christmasTree(height):
  # height: 聖誕樹的高度
  for i in range(height):
    # 使用迴圈控制每一行
    for j in range(height - i - 1):
      # 印出空白
      print(' ', end='')
    for k in range(2 * i + 1):
      # 印出星星
      print('*', end='')
    # 換行
    print()

  # 印出樹幹
  for i in range(height // 2):
    print(' ' * (height - 1) + '|')

# 聖誕樹要多高久把數字改成多少
christmasTree(6)

"""
聖誕節快樂 MERRY CHRISTMAS 
2022 / 12 / 25 by.KageRyo@codeRyo
"""
