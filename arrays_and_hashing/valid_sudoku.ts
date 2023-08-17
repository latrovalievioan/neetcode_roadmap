//https://leetcode.com/problems/valid-sudoku/

const isValidSudoku = (board: string[][]): boolean => {
  for (let row = 0; row < board.length; row++) {
    const set = new Set()
    for (let col = 0; col < board[row].length; col++) {
      const current = board[row][col]
      if (current === '.') continue
      if (set.has(current)) return false
      set.add(current)
    }
  }

  for (let col = 0; col < board[0].length; col++) {
    const set = new Set()
    for (let row = 0; row < board.length; row++) {
      const current = board[row][col]
      if (current === '.') continue
      if (set.has(current)) return false
      set.add(current)
    }
  }

  for (let rowOffset = 0; rowOffset < board.length; rowOffset += 3) {
    for (let colOffset = 0; colOffset < board[0].length; colOffset += 3) {
      const set = new Set()
      for (let r = rowOffset; r < rowOffset + 3; r++) {
        for (let c = colOffset; c < colOffset + 3; c++) {
          const current = board[r][c]
          if (current === '.') continue
          if (set.has(current)) return false
          set.add(current)
        }
      }
    }

  }
  return true
};


const test1 = [
  ["5", "3", ".", ".", "7", ".", ".", ".", "."],
  ["6", ".", ".", "1", "9", "5", ".", ".", "."],
  [".", "9", "8", ".", ".", ".", ".", "6", "."],
  ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
  ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
  ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
  [".", "6", ".", ".", ".", ".", "2", "8", "."],
  [".", ".", ".", "4", "1", "9", ".", ".", "5"],
  [".", ".", ".", ".", "8", ".", ".", "7", "9"]
]

const test2 = [
  ["8", "3", ".", ".", "7", ".", ".", ".", "."],
  ["6", ".", ".", "1", "9", "5", ".", ".", "."],
  [".", "9", "8", ".", ".", ".", ".", "6", "."],
  ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
  ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
  ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
  [".", "6", ".", ".", ".", ".", "2", "8", "."],
  [".", ".", ".", "4", "1", "9", ".", ".", "5"],
  [".", ".", ".", ".", "8", ".", ".", "7", "9"]
]

console.log(isValidSudoku((test2)))
