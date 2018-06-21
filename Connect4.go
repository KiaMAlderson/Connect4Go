package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
  "github.com/ahmetalpbalkan/go-cursor"
  "github.com/fatih/color"
)

//Alias for ease of use
var p = fmt.Print

//The board has 6 rows and 7 columns
type Board [6][7]string

func populateBoard(board Board) Board{
  //Board starting config : {"1", "2", "3", "4", "5", "6", "7"}
  //Columns labelled initially, then overwritten by play
  for i := 0; i < 6; i++ {
    for j := 0; j < 7; j++ {
      board[i][j] = strconv.Itoa(j+1)
    }
  }
  return board
}

func drawBoard(board Board) {
  //Prints board in a grid function showing most recent state of play
  var hor = " ----------------------------- \n"
  p(hor)
  for i := 0; i < 6; i++ {
    for j := 0; j < 7; j++ {
      toPrint := board[i][j]
      p(" | ")
      //Prints in either Red or Yellow for increased visability
      if (toPrint == "X") {
        x := color.New(color.FgRed)
        x.Print(toPrint)
      } else if (toPrint == "O"){
        o := color.New(color.FgYellow)
        o.Print(toPrint)
      } else {
        p(toPrint)
      }
    }
    p(" | \n")
    p(hor)
  }
}

func makeMove(turn string, board Board) Board{
  //Gets wanted column from the current Player
  //simulates dropping to the bottom of the available space
  //will reject if column is full or invalid

  //Flag to check if column is full
  var flag bool = false

  for flag == false {

    //Accepts col to place piece from user
    var column int64
    reader := bufio.NewReader(os.Stdin)
    p("Please enter a column from 1 to 7 that is available: ")
    col, _ := reader.ReadString('\n')
    col = strings.Trim(col, "\n")
    column, _ = strconv.ParseInt(col, 0, 64)

    for (column <= 0 || column > 7) {
      //Prevents illegal input
      p("Please enter a column from 1 to 7 that is available: ")
      col, _ := reader.ReadString('\n')
      col = strings.Trim(col, "\n")
      column, _ = strconv.ParseInt(col, 0, 64)
    }

    for (board[0][column - 1] == "X") || (board[0][column - 1] == "O") {
      //Checks if column is full, rejects if so
      p("Please enter a column from 1 to 7 that is available: ")
      col, _ := reader.ReadString('\n')
      col = strings.Trim(col, "\n")
      column, _ = strconv.ParseInt(col, 0, 64)
    }

    for i := 5; i >= 0; i-- {
      //Simulates dropping the piece down to the bottom of empty space in col
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

  //Checks all rows for a line of 4
  for i := 0; i < 6; i++ {
    for j := 0; j < 4; j++ {
      if (board[i][j] == board[i][j+1]) && (board[i][j] == board[i][j+2]) && (board[i][j] == board[i][j+3]){
        flag = true
      }
    }
  }


  //Checks all columns for a line of 4
  for i := 0; i < 3; i++ {
    for j := 0; j < 7; j++ {
      if (board[i][j] == board[i+1][j]) && (board[i][j] == board[i+2][j]) &&
         (board[i][j] == board[i+3][j]) && (board[i][j] == "X" || board[i][j] == "0"){
        flag = true
      }
    }
  }


  //Checks L-R diagonals for a line of 4
  for i := 0; i < 3; i++ {
    for j := 0; j < 4; j++ {
      if (board[i][j] == board[i+1][j+1]) && (board[i][j] == board[i+2][j+2]) && (board[i][j] == board[i+3][j+3]){
        flag = true
      }
    }
  }

  //Checks R-L diagonals for a line of 4
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
    //Convenience function for checking if a number is even
    return number%2 == 0
}

func main() {

  //Declare an instance of Board and initialise with starting values
  var board Board
  var turn string
  board = populateBoard(board)

  //Clears the terminal before play
  p(cursor.ClearEntireScreen())
  p(cursor.MoveTo(0, 0))

  p("-----------------------------------------------------------\n")
  color.Red("Welcome to Connect4!\n")
  //Gathers names from players
  reader := bufio.NewReader(os.Stdin)
  p("X Player 1 name : ")
  name1, _ := reader.ReadString('\n')
  p("O Player 2 name : ")
  name2, _ := reader.ReadString('\n')

  name1 = strings.Trim(name1, "\n")
  name2 = strings.Trim(name2, "\n")

  //Begin game loop. Max of 42 turns
  for i := 0; i < 42; i++ {
    p(cursor.ClearEntireScreen())
    p(cursor.MoveTo(0, 0))
    if Even(i){
      p("\n" + name1 + "'s turn!\n")
      turn = "X"
    } else {
      p("\n" + name2 + "'s turn!\n")
      turn = "O"
    }

    //Draw board again before each new move accepted
    drawBoard(board)
    board = makeMove(turn, board)

    //Check for win state, then exit game loop if true
    if checkStatus(board) {
      drawBoard(board)
      if Even(i){
        p("\n" + name1 + " wins!\n\n")
      } else {
        p("\n" + name2 + " wins!\n\n")
      }
      //Exit game successfully
      os.Exit(0)
    }

  }
  //Indicates a draw, before finishing game
  p("It's a draw!\n")
  p("Thanks for playing!\n")
}
