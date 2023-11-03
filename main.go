package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
		myApp := app.New()
		myWindow := myApp.NewWindow("Caro Game")
	
		boardSize := 15
		board := make([][]*canvas.Image, boardSize)
	
		// Khởi tạo người chơi hiện tại
		currentPlayer := "X"
	
		// Tải hình ảnh "X" và "O" từ tệp hoặc tài nguyên
		imageX, _ := canvas.NewImageFromResource(resourceX)
		imageO, _ := canvas.NewImageFromResource(resourceO)
	
		boardContainer := container.NewGridWithColumns(boardSize)
		for i := 0; i < boardSize; i++ {
			for j := 0; j < boardSize; j++ {
				board[i][j] = canvas.NewImageFromImage(imageX.Image)
				board[i][j].SetMinSize(boardCellSize)
				board[i][j].SetOnTapped(func() {
					handleCellTapped(board[i][j], &currentPlayer, imageX, imageO)
				})
				boardContainer.Add(board[i][j])
			}
		}
	
		myWindow.SetContent(container.NewVBox(
			widget.NewLabel("Caro Game"),
			boardContainer,
		))	
		myWindow.Resize(boardCellSize.Width*boardSize, boardCellSize.Height*boardSize)
		myWindow.ShowAndRun()
	}
	
func handleCellTapped(cell *canvas.Image, currentPlayer *string, imageX, imageO *canvas.Image) {
		if cell.Image == imageX.Image || cell.Image == imageO.Image {
			// Ô cờ đã được đánh, không làm gì cả
			return
		}
	
		if *currentPlayer == "X" {
			cell.Image = imageX.Image
		} else {
			cell.Image = imageO.Image
		}
	
		// Đổi người chơi
		if *currentPlayer == "X" {
			*currentPlayer = "O"
		} else {
			*currentPlayer = "X"
		}
	}
	
var (
	boardCellSize = fyne.NewSize(40, 40)
)
func loadImageResourceFromRelativePath(relativePath string) *fyne.StaticResource {
    resource := fyne.NewStaticResourceFromFile(relativePath)
    return resource
}
// Tải hình ảnh "X" và "O" từ tài nguyên
resourceX := loadImageResourceFromRelativePath("resources/X.png")
resourceO := loadImageResourceFromRelativePath("resources/O.png")

func handleCellTapped(cell *canvas.Text) {
	if cell.Text == "" {
		cell.Text = currentPlayer
		// Kiểm tra chiến thắng sau mỗi lượt đi
		if checkWin(cell.Text) {
			showWinMessage(currentPlayer)
		} else {
			// Đổi lượt chơi
			currentPlayer = switchPlayer(currentPlayer)
		}
	}
}
var (
	boardColor    = 0xffaaaaaa
	boardCellSize = fyne.NewSize(40, 40)
)
func checkWin(player string) bool {
	// Kiểm tra chiến thắng theo các hướng: ngang, dọc, chéo
	return checkHorizontal(player) || checkVertical(player) || checkDiagonal(player)
}

func checkHorizontal(player string) bool {
	// Kiểm tra chiến thắng theo hướng ngang
	for i := 0; i < boardSize; i++ {
		for j := 0; j <= boardSize-5; j++ {
			if board[i][j].Text == player && board[i][j+1].Text == player && board[i][j+2].Text == player &&
				board[i][j+3].Text == player && board[i][j+4].Text == player {
				return true
			}
		}
	}
	return false
}

func checkVertical(player string) bool {
	// Kiểm tra chiến thắng theo hướng dọc
	for i := 0; i <= boardSize-5; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j].Text == player && board[i+1][j].Text == player && board[i+2][j].Text == player &&
				board[i+3][j].Text == player && board[i+4][j].Text == player {
				return true
			}
		}
	}
	return false
}

func checkDiagonal(player string) bool {
	// Kiểm tra chiến thắng theo hướng chéo
	for i := 0; i <= boardSize-5; i++ {
		for j := 0; j <= boardSize-5; j++ {
			if board[i][j].Text == player && board[i+1][j+1].Text == player && board[i+2][j+2].Text == player &&
				board[i+3][j+3].Text == player && board[i+4][j+4].Text == player {
				return true
			}
		}
	}

	for i := 0; i <= boardSize-5; i++ {
		for j := 4; j < boardSize; j++ {
			if board[i][j].Text == player && board[i+1][j-1].Text == player && board[i+2][j-2].Text == player &&
				board[i+3][j-3].Text == player && board[i+4][j-4].Text == player {
				return true
			}
		}
	}
	return false
}
