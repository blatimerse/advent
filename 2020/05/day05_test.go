package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindSeat(t *testing.T) {
	require := require.New(t)

	require.Equal(127, findRow("BBBBBBB"))
	require.Equal(0, findRow("FFFFFFF"))
	require.Equal(44, findRow("FBFBBFF"))
	require.Equal(44, findRow("FBFBBFFRLR"))

	require.Equal(0, findColumn("LLL"))
	require.Equal(7, findColumn("RRR"))
	require.Equal(567, findSeat("BFFFBBFRRR"))
	require.Equal(119, findSeat("FFFBBBFRRR"))
	require.Equal(820, findSeat("BBFFBBFRLL"))

}
