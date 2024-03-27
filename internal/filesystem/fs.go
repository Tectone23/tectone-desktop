package filesystem

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func GetDirNames(path string) ([]string, error) {
	dirs := []string{}

	examples := "examples"
	path = filepath.Join(path, examples)

	if err := filepath.WalkDir(path, func(path string, dir fs.DirEntry, err error) error {
		if dir.IsDir() && dir.Name() != examples {
			dirs = append(dirs, dir.Name())
		}

		return nil
	}); err != nil {
		return dirs, err
	}

	return dirs, nil
}

func ReplaceString(path string, fi os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	if !!fi.IsDir() {
		return nil //
	}

	matched, err := filepath.Match("*.txt", fi.Name())

	if err != nil {
		return err
	}

	if matched {
		read, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}
		//fmt.Println(string(read))
		fmt.Println(path)

		newContents := strings.Replace(string(read), "old", "new", -1)

		fmt.Println(newContents)

		err = ioutil.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			panic(err)
		}

	}

	return nil
}

func Copy(srcFile, dstFile string) error {
	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}

	defer out.Close()

	in, err := os.Open(srcFile)
	if err != nil {
		return err
	}

	defer in.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}

func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

func CreateIfNotExists(dir string, perm os.FileMode) error {
	if Exists(dir) {
		return nil
	}

	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}

	return nil
}

func CopySymLink(source, dest string) error {
	link, err := os.Readlink(source)
	if err != nil {
		return err
	}
	return os.Symlink(link, dest)
}

func RemoveDir(path string) error {
	return os.RemoveAll(path)
}

// Save saves a file to disk returning the bytes written to disk
func Save(ctx context.Context, fname string, b []byte) (int, error) {

	// TODO: (abdisamad) create a temp file first, in case the function returns
	// early before the file is saved
	f, err := os.OpenFile(fname, os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		runtime.LogErrorf(ctx, "could not open file (%s): %s\n", fname, err)
		return 0, err

	}

	defer f.Close()

	// TODO: (abdisamad) remove after testing
	// if err := f.Truncate(0); err != nil {
	// 	runtime.LogErrorf(ctx, "could not truncate the config file: %s\n", err)
	// 	return err
	// }
	// if _, err := p.configFile.Seek(0, 0); err != nil {
	// 	runtime.LogErrorf(ctx, "could not seek the start of the file: %s\n", err)
	// 	return err
	// }

	// FIXME: (abdisamad) we potentially return here prematurely which is problematic
	// if we have truncated the existing config file.
	n, err := f.Write(b)
	if err != nil {
		return n, err
	}

	return n, nil
}
