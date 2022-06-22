//go:generate fyne bundle -o bundled-pieces.go pieces

package main

import (
	"fyne.io/fyne/v2"
	"github.com/notnil/chess"
)

func resourceForPiece(p chess.Piece) fyne.Resource {
	switch p.Color() {
	case chess.Black:
		switch p.Type() {
		case chess.Pawn:
			return resourceBlackPawnSvg
		case chess.Bishop:
			return resourceBlackBishopSvg
		case chess.King:
			return resourceBlackKingSvg
		case chess.Knight:
			return resourceBlackKnightSvg
		case chess.Queen:
			return resourceBlackQueenSvg
		case chess.Rook:
			return resourceBlackRookSvg
		}
	case chess.White:
		switch p.Type() {
		case chess.Pawn:
			return resourceWhitePawnSvg
		case chess.Bishop:
			return resourceWhiteBishopSvg
		case chess.King:
			return resourceWhiteKingSvg
		case chess.Knight:
			return resourceWhiteKnightSvg
		case chess.Queen:
			return resourceWhiteQueenSvg
		case chess.Rook:
			return resourceWhiteRookSvg
		}
	}
	return nil
}
