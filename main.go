// The purpose of this utility is to easily test the library from the command line.
package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/Preciselyco/go-diff/diffmatchpatch"
	"github.com/alecthomas/kong"
)

func readFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

type DiffCmd struct {
	CleanupEfficiency       bool    `help:"Filter output through CleanupEfficiency()"`
	CleanupMerge            bool    `help:"Filter output through CleanupMerge()"`
	CleanupSemantic         bool    `help:"Filter output through CleanupSemantic()"`
	CleanupSemanticLossless bool    `help:"Filter output through CleanupSemanticLossless()"`
	DiffTimeout             int     `help:"Input to DiffMatchPatch.DiffTimeout (seconds)"`
	DiffEditCost            int     `help:"Input to DiffMatchPatch.DiffEditCost"`
	MatchDistance           int     `help:"Input to DiffMatchPatch.MatchDistance"`
	PatchDeleteThreshold    float64 `help:"Input to DiffMatchPatch.PatchDeleteThreshold"`
	PatchMargin             int     `help:"Input to DiffMatchPatch.PatchMargin"`
	MatchThreshold          float64 `help:"Input to DiffMatchPatch.MatchThreshold"`
	CheckLines              bool    `help:"Input to DiffMatchPatch.DiffMain()"`
	OldFile                 string  `arg required help:"Original file"`
	NewFile                 string  `arg required help:"File which is compared to original file"`
}

func (cmd *DiffCmd) Run() error {
	oldString := readFile(cmd.OldFile)
	newString := readFile(cmd.NewFile)

	dmp := diffmatchpatch.New()
	if cmd.DiffTimeout > 0 {
		dmp.DiffTimeout = time.Second * time.Duration(cmd.DiffTimeout)
	}
	if cmd.DiffEditCost > 0 {
		dmp.DiffEditCost = cmd.DiffEditCost
	}
	if cmd.MatchDistance > 0 {
		dmp.MatchDistance = cmd.MatchDistance
	}
	if cmd.PatchDeleteThreshold > 0.0 {
		dmp.PatchDeleteThreshold = cmd.PatchDeleteThreshold
	}
	if cmd.PatchMargin > 0 {
		dmp.PatchMargin = cmd.PatchMargin
	}
	if cmd.MatchThreshold > 0.0 {
		dmp.MatchThreshold = cmd.MatchThreshold
	}
	diffs := dmp.DiffMain(oldString, newString, cmd.CheckLines)

	if cmd.CleanupEfficiency {
		diffs = dmp.DiffCleanupEfficiency(diffs)
	}
	if cmd.CleanupMerge {
		diffs = dmp.DiffCleanupMerge(diffs)
	}
	if cmd.CleanupSemantic {
		diffs = dmp.DiffCleanupSemantic(diffs)
	}
	if cmd.CleanupSemanticLossless {
		diffs = dmp.DiffCleanupSemanticLossless(diffs)
	}
	fmt.Println(dmp.DiffPrettyText(diffs))

	return nil
}

func main() {
	cli := struct {
		Diff DiffCmd `cmd help:"Print the diff of two files"`
	}{}
	ctx := kong.Parse(&cli)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
