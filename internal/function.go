package internal

import (
	"embed"
	"fmt"
	"io"
	"os"
	"strings"
)

var templateDir = "template/"

type Project struct {
	ModuleName string
	Directory  string
	Template   embed.FS
}

func (p Project) Start() error {

	err := p.InitializeProject()
	if err != nil {
		return err
	}

	err = p.CreateEntity()
	if err != nil {
		return err
	}

	err = p.CreateRepository()
	if err != nil {
		return err
	}

	err = p.CreateConstant()
	if err != nil {
		return err
	}

	err = p.CreateUtil()
	if err != nil {
		return err
	}

	err = p.CreateLogger()
	if err != nil {
		return err
	}

	err = p.CreateModel()
	if err != nil {
		return err
	}

	err = p.CreateConverter()
	if err != nil {
		return err
	}

	err = p.CreateUseCase()
	if err != nil {
		return err
	}

	err = p.CreateDelivery()
	if err != nil {
		return err
	}

	err = p.CreateConfig()
	if err != nil {
		return err
	}

	err = p.CreateEntryPoint()
	if err != nil {
		return err
	}

	err = p.CreateConfigFile()
	if err != nil {
		return err
	}

	return nil
}

func (p Project) InitializeProject() error {

	err := os.Mkdir(p.Directory, 0755)
	if err != nil {
		return err
	}
	err = os.Mkdir(p.Directory+"/internal", 0755)
	if err != nil {
		return err
	}

	templateMod := templateDir + "go_mod.file"
	destinationMode := fmt.Sprintf("%s/go.mod", p.Directory)
	srcFile, destFile, err := p.copyFile(templateMod, destinationMode)
	defer srcFile.Close()
	defer destFile.Close()
	if err != nil {
		return err
	}

	err = modifyAndCopy(srcFile, destFile, "__module_name__", p.ModuleName)
	if err != nil {
		fmt.Println("Gagal menyalin dan memodifikasi file:", err)
		return err
	}

	return nil
}

func (p Project) CreateEntity() error {

	err := os.Mkdir(p.Directory+"/internal/entity", 0755)
	if err != nil {
		return err
	}

	templateMod := templateDir + "internal/entity/entity.file"
	destinationMode := fmt.Sprintf("%s/internal/entity/entity.go", p.Directory)
	srcFile, destFile, err := p.copyFile(templateMod, destinationMode)
	defer srcFile.Close()
	defer destFile.Close()
	if err != nil {
		return err
	}

	// Menyalin konten file asli ke file tujuan
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		fmt.Println("Gagal menyalin file entity:", err)
		return err
	}

	return nil
}

func (p Project) CreateRepository() error {

	err := os.Mkdir(p.Directory+"/internal/repository", 0755)
	if err != nil {
		return err
	}

	templateMod := templateDir + "internal/repository/sample_repository.file"
	destinationMode := fmt.Sprintf("%s/internal/repository/sample_repository.go", p.Directory)
	srcFile, destFile, err := p.copyFile(templateMod, destinationMode)
	defer srcFile.Close()
	defer destFile.Close()
	if err != nil {
		return err
	}

	err = modifyAndCopy(srcFile, destFile, "__module_name__", p.ModuleName)
	if err != nil {
		fmt.Println("Gagal menyalin dan memodifikasi file:", err)
		return err
	}

	templateModImpl := templateDir + "internal/repository/sample_repository_impl.file"
	destinationModeImpl := fmt.Sprintf("%s/internal/repository/sample_repository_impl.go", p.Directory)
	srcFileImpl, destFileImpl, err := p.copyFile(templateModImpl, destinationModeImpl)
	defer srcFileImpl.Close()
	defer destFileImpl.Close()
	if err != nil {
		return err
	}

	err = modifyAndCopy(srcFileImpl, destFileImpl, "__module_name__", p.ModuleName)
	if err != nil {
		fmt.Println("Gagal menyalin dan memodifikasi file:", err)
		return err
	}

	return nil
}

func (p Project) CreateConstant() error {

	err := os.Mkdir(p.Directory+"/internal/constants", 0755)
	if err != nil {
		return err
	}

	templateMod := templateDir + "internal/constants/constants.file"
	destinationMode := fmt.Sprintf("%s/internal/constants/constants.go", p.Directory)
	srcFile, destFile, err := p.copyFile(templateMod, destinationMode)
	defer srcFile.Close()
	defer destFile.Close()
	if err != nil {
		return err
	}

	// Menyalin konten file asli ke file tujuan
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		fmt.Println("Gagal menyalin file constants:", err)
		return err
	}

	return nil
}

func (p Project) CreateUtil() error {

	err := os.Mkdir(p.Directory+"/internal/util", 0755)
	if err != nil {
		return err
	}

	templateMod := templateDir + "internal/util/common.file"
	destinationMode := fmt.Sprintf("%s/internal/util/common.go", p.Directory)
	srcFile, destFile, err := p.copyFile(templateMod, destinationMode)
	defer srcFile.Close()
	defer destFile.Close()
	if err != nil {
		return err
	}

	err = modifyAndCopy(srcFile, destFile, "__module_name__", p.ModuleName)
	if err != nil {
		fmt.Println("Gagal menyalin dan memodifikasi file:", err)
		return err
	}
	return nil
}

func (p Project) CreateLogger() error {

	err := os.Mkdir(p.Directory+"/internal/logger", 0755)
	if err != nil {
		return err
	}

	templateMod := templateDir + "internal/logger/logger.file"
	destinationMode := fmt.Sprintf("%s/internal/logger/logger.go", p.Directory)
	srcFile, destFile, err := p.copyFile(templateMod, destinationMode)
	defer srcFile.Close()
	defer destFile.Close()
	if err != nil {
		return err
	}

	err = modifyAndCopy(srcFile, destFile, "__module_name__", p.ModuleName)
	if err != nil {
		fmt.Println("Gagal menyalin dan memodifikasi file:", err)
		return err
	}
	return nil
}

func (p Project) CreateModel() error {

	err := os.Mkdir(p.Directory+"/internal/model/", 0755)
	if err != nil {
		return err
	}
	err = os.Mkdir(p.Directory+"/internal/model/dto", 0755)
	if err != nil {
		return err
	}

	templateMod := templateDir + "internal/model/dto/common.file"
	destinationMode := fmt.Sprintf("%s/internal/model/dto/common.go", p.Directory)
	srcFile, destFile, err := p.copyFile(templateMod, destinationMode)
	defer srcFile.Close()
	defer destFile.Close()
	if err != nil {
		return err
	}

	// Menyalin konten file asli ke file tujuan
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		fmt.Println("Gagal menyalin file:", err)
		return err
	}

	templateModSample := templateDir + "internal/model/dto/sample.file"
	destinationModeSample := fmt.Sprintf("%s/internal/model/dto/sample.go", p.Directory)
	srcFileSample, destFileSample, err := p.copyFile(templateModSample, destinationModeSample)
	defer srcFileSample.Close()
	defer destFileSample.Close()
	if err != nil {
		return err
	}

	// Menyalin konten file asli ke file tujuan
	_, err = io.Copy(destFileSample, srcFileSample)
	if err != nil {
		fmt.Println("Gagal menyalin file:", err)
		return err
	}

	return nil
}

func (p Project) CreateConverter() error {

	err := os.Mkdir(p.Directory+"/internal/converter", 0755)
	if err != nil {
		return err
	}

	templateMod := templateDir + "internal/converter/entity.file"
	destinationMode := fmt.Sprintf("%s/internal/converter/entity.go", p.Directory)
	srcFile, destFile, err := p.copyFile(templateMod, destinationMode)
	defer srcFile.Close()
	defer destFile.Close()
	if err != nil {
		return err
	}

	err = modifyAndCopy(srcFile, destFile, "__module_name__", p.ModuleName)
	if err != nil {
		fmt.Println("Gagal menyalin dan memodifikasi file:", err)
		return err
	}
	return nil
}

func (p Project) CreateUseCase() error {

	err := os.Mkdir(p.Directory+"/internal/usecase", 0755)
	if err != nil {
		return err
	}

	templateMod := templateDir + "internal/usecase/sample_usecase.file"
	destinationMode := fmt.Sprintf("%s/internal/usecase/sample_usecase.go", p.Directory)
	srcFile, destFile, err := p.copyFile(templateMod, destinationMode)
	defer srcFile.Close()
	defer destFile.Close()
	if err != nil {
		return err
	}

	err = modifyAndCopy(srcFile, destFile, "__module_name__", p.ModuleName)
	if err != nil {
		fmt.Println("Gagal menyalin dan memodifikasi file:", err)
		return err
	}

	templateModImpl := templateDir + "internal/usecase/sample_usecase_impl.file"
	destinationModeImpl := fmt.Sprintf("%s/internal/usecase/sample_usecase_impl.go", p.Directory)
	srcFileImpl, destFileImpl, err := p.copyFile(templateModImpl, destinationModeImpl)
	defer srcFileImpl.Close()
	defer destFileImpl.Close()
	if err != nil {
		return err
	}

	err = modifyAndCopy(srcFileImpl, destFileImpl, "__module_name__", p.ModuleName)
	if err != nil {
		fmt.Println("Gagal menyalin dan memodifikasi file:", err)
		return err
	}

	return nil
}

func (p Project) CreateDelivery() error {

	err := os.Mkdir(p.Directory+"/internal/delivery", 0755)
	if err != nil {
		return err
	}

	err = os.Mkdir(p.Directory+"/internal/delivery/http", 0755)
	if err != nil {
		return err
	}

	err = os.Mkdir(p.Directory+"/internal/delivery/http/handler", 0755)
	if err != nil {
		return err
	}

	err = os.Mkdir(p.Directory+"/internal/delivery/http/middleware", 0755)
	if err != nil {
		return err
	}

	err = os.Mkdir(p.Directory+"/internal/delivery/http/route", 0755)
	if err != nil {
		return err
	}

	templateMod := templateDir + "internal/delivery/http/handler/sample_handler.file"
	destinationMode := fmt.Sprintf("%s/internal/delivery/http/handler/sample_handler.go", p.Directory)
	srcFile, destFile, err := p.copyFile(templateMod, destinationMode)
	defer srcFile.Close()
	defer destFile.Close()
	if err != nil {
		return err
	}

	err = modifyAndCopy(srcFile, destFile, "__module_name__", p.ModuleName)
	if err != nil {
		fmt.Println("Gagal menyalin dan memodifikasi file:", err)
		return err
	}

	templateModMiddleware := templateDir + "internal/delivery/http/middleware/authentication.file"
	destinationModeMiddleware := fmt.Sprintf("%s/internal/delivery/http/middleware/authentication.go", p.Directory)
	srcFileMiddleware, destFileMiddleware, err := p.copyFile(templateModMiddleware, destinationModeMiddleware)
	defer srcFileMiddleware.Close()
	defer destFileMiddleware.Close()
	if err != nil {
		return err
	}

	// Menyalin konten file asli ke file tujuan
	_, err = io.Copy(destFileMiddleware, srcFileMiddleware)
	if err != nil {
		fmt.Println("Gagal menyalin file:", err)
		return err
	}

	templateModMiddleware2 := templateDir + "internal/delivery/http/middleware/common.file"
	destinationModeMiddleware2 := fmt.Sprintf("%s/internal/delivery/http/middleware/common.go", p.Directory)
	srcFileMiddleware2, destFileMiddleware2, err := p.copyFile(templateModMiddleware2, destinationModeMiddleware2)
	defer srcFileMiddleware2.Close()
	defer destFileMiddleware2.Close()
	if err != nil {
		return err
	}

	err = modifyAndCopy(srcFileMiddleware2, destFileMiddleware2, "__module_name__", p.ModuleName)
	if err != nil {
		fmt.Println("Gagal menyalin dan memodifikasi file:", err)
		return err
	}

	templateModRoute := templateDir + "internal/delivery/http/route/route.file"
	destinationModeRoute := fmt.Sprintf("%s/internal/delivery/http/route/route.go", p.Directory)
	srcFileRoute, destFileRoute, err := p.copyFile(templateModRoute, destinationModeRoute)
	defer srcFileRoute.Close()
	defer destFileRoute.Close()
	if err != nil {
		return err
	}

	err = modifyAndCopy(srcFileRoute, destFileRoute, "__module_name__", p.ModuleName)
	if err != nil {
		fmt.Println("Gagal menyalin dan memodifikasi file:", err)
		return err
	}

	return nil
}

func (p Project) CreateConfig() error {
	err := os.Mkdir(p.Directory+"/internal/config", 0755)
	if err != nil {
		return err
	}

	templateMod := templateDir + "internal/config/app.file"
	destinationMode := fmt.Sprintf("%s/internal/config/app.go", p.Directory)
	srcFile, destFile, err := p.copyFile(templateMod, destinationMode)
	defer srcFile.Close()
	defer destFile.Close()
	if err != nil {
		return err
	}

	err = modifyAndCopy(srcFile, destFile, "__module_name__", p.ModuleName)
	if err != nil {
		fmt.Println("Gagal menyalin dan memodifikasi file:", err)
		return err
	}

	templateModDb := templateDir + "internal/config/db.file"
	destinationModeDb := fmt.Sprintf("%s/internal/config/db.go", p.Directory)
	srcFileDb, destFileDb, err := p.copyFile(templateModDb, destinationModeDb)
	defer srcFileDb.Close()
	defer destFileDb.Close()
	if err != nil {
		return err
	}

	// Menyalin konten file asli ke file tujuan
	_, err = io.Copy(destFileDb, srcFileDb)
	if err != nil {
		fmt.Println("Gagal menyalin file:", err)
		return err
	}

	templateModFiber := templateDir + "internal/config/fiber.file"
	destinationModeFiber := fmt.Sprintf("%s/internal/config/fiber.go", p.Directory)
	srcFileFiber, destFileFiber, err := p.copyFile(templateModFiber, destinationModeFiber)
	defer srcFileFiber.Close()
	defer destFileFiber.Close()
	if err != nil {
		return err
	}

	err = modifyAndCopy(srcFileFiber, destFileFiber, "__module_name__", p.ModuleName)
	if err != nil {
		fmt.Println("Gagal menyalin dan memodifikasi file:", err)
		return err
	}

	templateModValidator := templateDir + "internal/config/validator.file"
	destinationModeValidator := fmt.Sprintf("%s/internal/config/validator.go", p.Directory)
	srcFileValidator, destFileValidator, err := p.copyFile(templateModValidator, destinationModeValidator)
	defer srcFileValidator.Close()
	defer destFileValidator.Close()
	if err != nil {
		return err
	}

	// Menyalin konten file asli ke file tujuan
	_, err = io.Copy(destFileValidator, srcFileValidator)
	if err != nil {
		fmt.Println("Gagal menyalin file:", err)
		return err
	}

	templateModViper := templateDir + "internal/config/viper.file"
	destinationModeViper := fmt.Sprintf("%s/internal/config/viper.go", p.Directory)
	srcFileViper, destFileViper, err := p.copyFile(templateModViper, destinationModeViper)
	defer srcFileViper.Close()
	defer destFileViper.Close()
	if err != nil {
		return err
	}

	// Menyalin konten file asli ke file tujuan
	_, err = io.Copy(destFileViper, srcFileViper)
	if err != nil {
		fmt.Println("Gagal menyalin file:", err)
		return err
	}

	return nil

}

func (p Project) CreateEntryPoint() error {
	err := os.Mkdir(p.Directory+"/cmd", 0755)
	if err != nil {
		return err
	}
	err = os.Mkdir(p.Directory+"/cmd/rest", 0755)
	if err != nil {
		return err
	}

	templateMod := templateDir + "cmd/rest/main.file"
	destinationMode := fmt.Sprintf("%s/cmd/rest/main.go", p.Directory)
	srcFile, destFile, err := p.copyFile(templateMod, destinationMode)
	defer srcFile.Close()
	defer destFile.Close()
	if err != nil {
		return err
	}

	err = modifyAndCopy(srcFile, destFile, "__module_name__", p.ModuleName)
	if err != nil {
		fmt.Println("Gagal menyalin dan memodifikasi file:", err)
		return err
	}
	return nil
}

func (p Project) CreateConfigFile() error {
	templateMod := templateDir + "config.json"
	destinationMode := fmt.Sprintf("%s/config.json", p.Directory)
	srcFile, destFile, err := p.copyFile(templateMod, destinationMode)
	defer srcFile.Close()
	defer destFile.Close()
	if err != nil {
		return err
	}

	// Menyalin konten file asli ke file tujuan
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		fmt.Println("Gagal menyalin file:", err)
		return err
	}
	return nil
}

func (p Project) copyFile(source, destination string) (*os.File, *os.File, error) {
	// Membuka file asli
	data, err := p.Template.ReadFile(source)
	if err != nil {
		panic(err)
	}
	// Membuat file sementara di filesystem
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

	// Menulis data ke file sementara
	if _, err := tmpfile.Write(data); err != nil {
		panic(err)
	}

	srcFile, err := os.Open(tmpfile.Name())
	if err != nil {
		return nil, nil, err
	}
	// Membuat file tujuan
	destFile, err := os.Create(destination)
	if err != nil {
		fmt.Println("Gagal membuat file tujuan:", err)
		return nil, nil, err
	}

	return srcFile, destFile, nil
}

func modifyAndCopy(src io.Reader, dest io.Writer, before, after string) error {
	// Buat buffer untuk membaca file asli
	buf := make([]byte, 1024)

	for {
		// Baca data dari file asli ke buffer
		n, err := src.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		// Ubah konten di dalam buffer
		content := string(buf[:n])
		modifiedContent := strings.ReplaceAll(content, before, after)

		// Tulis konten yang sudah dimodifikasi ke file tujuan
		_, err = dest.Write([]byte(modifiedContent))
		if err != nil {
			return err
		}
	}

	return nil
}
