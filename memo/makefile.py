def main(number):
  # 何番目まで作るかを number に入れる
  for i in range(1, number+1):
    # 2 → 02 に
    twoInt = str(i).zfill(2)
    filename = "chapter" + twoInt+".md"
    print("実行中:"+filename)
    # 文字コードをUTF-8, 改行コードをLFに
    with open(filename, "w", encoding="utf-8", newline="\n") as f:
      f.write("## 第"+twoInt+"章 〇〇をする"+"\n")

if __name__ == "__main__":
  main(8)

# 改行コードの確認
# $ od -c [ファイル名]
