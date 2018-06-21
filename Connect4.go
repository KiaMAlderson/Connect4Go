package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
)

//Aliase for ease of use
var p = fmt.Print

//The board has 6 rows and 7 columns
//Board starting config : {"1", "2", "3", "4", "5", "6", "7"}
//Done by using populateBoard()
type Board [6][7]string

func populateBoard(board Board) Board{
  for i := 0; i < 6; i++ {
    for j := 0; j < 7; j++ {
      board[i][j] = strconv.Itoa(j+1)
    }
  }

  return board
}

func drawBoard(board Board) {
  //6 rows, 7 columns
  var hor = " ----------------------------- \n"
  p(hor)
  for i := 0; i < 6; i++ {
    for j := 0; j < 7; j++ {
      p(" | ", board[i][j])
    }
    p(" | \n")
    p(hor)
  }
}

func makeMove(turn string, board Board) Board{
  //Gets wanted column from the current Player
  //simulates dropping to the bottom of the available space
  //will reject if column is full or invalid



  //flag to check if column is full
  var flag bool = false

  for flag == false {

    //Accepts col to place piece from user, then updates pieces
    var column int64
    reader := bufio.NewReader(os.Stdin)
    p("Please enter a column from 1 to 7 that is available: ")
    col, _ := reader.ReadString('\n')
    col = strings.Trim(col, "\n")
    column, _ = strconv.ParseInt(col, 0, 64)

    for (column <= 0 || column > 7) {
      p("Please enter a column from 1 to 7 that is available: ")
      col, _ := reader.ReadString('\n')
      col = strings.Trim(col, "\n")
      column, _ = strconv.ParseInt(col, 0, 64)
    }

    for (board[0][column - 1] == "X") || (board[0][column - 1] == "O") {
      p("Please enter a column from 1 to 7 that is available: ")
      col, _ := reader.ReadString('\n')
      col = strings.Trim(col, "\n")
      column, _ = strconv.ParseInt(col, 0, 64)
    }

    for i := 5; i >= 0; i-- {
      if (board[i][column - 1] == "X") || (board[i][column - 1] == "O") {
        continue
        } else {
          board[i][column - 1] = turn
          flag = true
          break
        }
      }
    }
    return board
}

func checkStatus(board Board) bool {
  //Checks for winning config of 4

  var flag bool = false

  //j - j+1 - j+2 - j+3
  // 0 <= j < 4 for rows
  for i := 0; i < 6; i++ {
    for j := 0; j < 4; j++ {
      if (board[i][j] == board[i][j+1]) && (board[i][j] == board[i][j+2]) && (board[i][j] == board[i][j+3]){
        flag = true
      }
    }
  }


  //i - i+1 - i+2 - i+3
  // 0 <= i < 3 for columns
  for i := 0; i < 3; i++ {
    for j := 0; j < 7; j++ {
      if (board[i][j] == board[i+1][j]) && (board[i][j] == board[i+2][j]) &&
         (board[i][j] == board[i+3][j]) && (board[i][j] == "X" || board[i][j] == "0"){
        flag = true
      }
    }
  }


  // (i, j) - (i+1, j+1) - (i+2, j+2) - (i+3, j+3)
  // (0, 0) <= (i, j) < (3, 4) for L-R diagonals \
  for i := 0; i < 3; i++ {
    for j := 0; j < 4; j++ {
      if (board[i][j] == board[i+1][j+1]) && (board[i][j] == board[i+2][j+2]) && (board[i][j] == board[i+3][j+3]){
        flag = true
      }
    }
  }

  // (i, j) - (i+1, j-1) - (i+2, j-2) - (i+3, j-3)
  // (0, 3) <= (i, j) < (3, 7) for R-L diagonals /
  for i := 0; i < 3; i++ {
    for j := 3; j < 7; j++ {
      if (board[i][j] == board[i+1][j-1]) && (board[i][j] == board[i+2][j-2]) && (board[i][j] == board[i+3][j-3]){
        flag = true
      }
    }
  }

  return flag
}

func Even(number int) bool {
    return number%2 == 0
}

func main() {

  //Declare an instance of Board and initialise with starting values
  var board Board
  var turn string
  board = populateBoard(board)

  p("-----------------------------------------------------------\n")
  p("Welcome to Connect4!\n")
  reader := bufio.NewReader(os.Stdin)
  p("X Player 1 name : ")
  name1, _ := reader.ReadString('\n')
  p("O Player 2 name : ")
  name2, _ := reader.ReadString('\n')

  name1 = strings.Trim(name1, "\n")
  name2 = strings.Trim(name2, "\n")

  for i := 0; i < 42; i++ {
    if Even(i){
      p("\n" + name1 + "'s turn!\n")
      turn = "X"
    } else {
      p("\n" + name2 + "'s turn!\n")
      turn = "O"
    }
    drawBoard(board)
    board = makeMove(turn, board)

    if checkStatus(board) {
      if Even(i){
        p("\n" + name1 + " wins!\n")
      } else {
        p("\n" + name2 + " wins!\n")
      }
      os.Exit(0)
    }

  }
  p("It's a draw!\n")
  p("Thanks for playing!\n")
}
