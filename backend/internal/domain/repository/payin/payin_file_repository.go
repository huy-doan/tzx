package repository

import (
	"context"

	"github.com/test-tzs/nomraeite/internal/datastructure/inputdata"
	model "github.com/test-tzs/nomraeite/internal/domain/model/payin"
	object "github.com/test-tzs/nomraeite/internal/domain/object/payin"
)

type PayinFileRepository interface {
	// Create a new PayinFile record
	Create(ctx context.Context, file *model.PayinFile) (*model.PayinFile, error)

	// UpdateStatus updates the status of a PayinFile record
	UpdateStatus(ctx context.Context, file *model.PayinFile) error

	// UpdateImportStatus updates the import status and has_data_record fields of a PayinFile record
	UpdateImportStatus(ctx context.Context, file *model.PayinFile) error

	// FindByFilename checks if a file exists by its filename and returns its PayinFile model
	FindByFilename(ctx context.Context, filename string) (*model.PayinFile, error)

	// GetByID retrieves a PayinFile record by its ID
	GetByID(ctx context.Context, id int) (*model.PayinFile, error)

	// GetPaginatedPayinFilesByDate retrieves payin files by target date with pagination
	GetPaginatedPayinFilesByDate(ctx context.Context, targetDate string, limit, offset int) ([]*model.PayinFile, error)

	// GetPaginatedPayinFilesByDateAndType retrieves payin files by target date and file type with pagination
	GetPaginatedPayinFilesByDateAndType(ctx context.Context, targetDate string, fileType object.PayinFileType, limit, offset int) ([]*model.PayinFile, error)

	ListPayinFiles(ctx context.Context, params *inputdata.PayinFileListInputData) (*model.PaginatedPayinFileResult, error)
}
