package service

import (
	"github.com/egasa21/si-lab-api-go/internal/dto"
	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/egasa21/si-lab-api-go/internal/pkg"
	"github.com/rs/zerolog/log"
)

type StudentDataService interface {
	GetStudentPracticumActivity(userID int) ([]dto.StudentPracticumActivity, error)
	GetStudentSchedules(studentID int) ([]dto.StudentSchedules, error)
}

type studentDataService struct {
	userPracticumCheckpointService UserPracticumCheckpointService
	practicumService               PracticumService
	practicumModuleService         PracticumModuleService
	practicumModuleContentService  PracticumModuleContentService
	studentClassEnrollmentService  StudentClassEnrollmentService
	practicumClassService          PracticumClassService
	userPracticumProgressService   UserPracticumProgressService
}

func NewStudentDataService(userPracticumCheckpointService UserPracticumCheckpointService, practicumService PracticumService, practicumModuleService PracticumModuleService, practicumModuleContentService PracticumModuleContentService, studentClasEstudentClassEnrollmentService StudentClassEnrollmentService, practicumClassService PracticumClassService, userPracticumProgressService UserPracticumProgressService) StudentDataService {
	return &studentDataService{
		userPracticumCheckpointService: userPracticumCheckpointService,
		practicumService:               practicumService,
		practicumModuleService:         practicumModuleService,
		practicumModuleContentService:  practicumModuleContentService,
		studentClassEnrollmentService:  studentClasEstudentClassEnrollmentService,
		practicumClassService:          practicumClassService,
		userPracticumProgressService:   userPracticumProgressService,
	}
}

func (s *studentDataService) GetStudentPracticumActivity(userID int) ([]dto.StudentPracticumActivity, error) {
	userCheckpoints, err := s.userPracticumCheckpointService.GetCheckpointByUser(userID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user practicum checkpoint")
		return nil, err
	}
	if userCheckpoints == nil {
		log.Info().Msg("Checkpoint not found")
	}

	// log.Info().Int("count", len(userCheckpoints)).Msgf("Successfully retrieved %d user practicum checkpoints for user %d", len(userCheckpoints), userID)

	practicumDataIDs := make(map[int]bool)
	practicumModuleDataIDs := make(map[int]bool)
	practicumModuleContentDataIDs := make(map[int]bool)

	for _, item := range userCheckpoints {
		practicumDataIDs[item.PracticumID] = true
		practicumModuleDataIDs[item.ModuleID] = true
		practicumModuleContentDataIDs[item.ContentID] = true
	}

	practicumIDs := pkg.GetKeysFromMap(practicumDataIDs)
	practicumModuleIDs := pkg.GetKeysFromMap(practicumModuleDataIDs)
	practicumModuleContentIDs := pkg.GetKeysFromMap(practicumModuleContentDataIDs)

	practicums, err := s.practicumService.GetPracticumByIDs(practicumIDs)
	if err != nil {
		return nil, err
	}

	practicumProgresses, err := s.userPracticumProgressService.GetProgressByPracticumIDs(practicumIDs)
	if err != nil {
		return nil, err
	}

	modules, err := s.practicumModuleService.GetModuleByIDs(practicumModuleIDs)
	if err != nil {
		return nil, err
	}

	moduleContents, err := s.practicumModuleContentService.GetContentByIDs(practicumModuleContentIDs)
	if err != nil {
		return nil, err
	}

	practicumMap := make(map[int]model.Practicum)
	for _, item := range practicums {
		practicumMap[item.ID] = item
	}

	practicumProgressMap := make(map[int]model.UserPracticumProgress)
	for _, item := range practicumProgresses {
		practicumProgressMap[item.ID] = item
	}

	moduleMap := make(map[int]model.PracticumModule)
	for _, item := range modules {
		moduleMap[item.ID] = item
	}

	contentMap := make(map[int]model.PracticumModuleContent)
	for _, item := range moduleContents {
		contentMap[item.IDContent] = item
	}

	practicumActivities := make([]dto.StudentPracticumActivity, len(userCheckpoints))
	for i, item := range userCheckpoints {
		practicumActivitiesObj := dto.StudentPracticumActivity{
			ID:                item.ID,
			PracticumName:     practicumMap[item.PracticumID].Name,
			PracticumProgress: int(practicumProgressMap[item.PracticumID].Progress),
			ModuleName:        moduleMap[item.ModuleID].Title,
			ModuleContentName: contentMap[item.ContentID].Title,
			ModuleSequence:    contentMap[item.ContentID].Sequence,
			ModuleContentID:   contentMap[item.ContentID].IDContent,
		}
		practicumActivities[i] = practicumActivitiesObj
	}

	return practicumActivities, nil

}

func (s *studentDataService) GetStudentSchedules(studentID int) ([]dto.StudentSchedules, error) {
	studentEnrollments, err := s.studentClassEnrollmentService.GetEnrollmentsByStudentID(studentID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get student enrollments")
		return nil, err
	}

	if studentEnrollments == nil {
		log.Info().Msg("Student enrollments not found")
	}

	log.Info().Int("count", len(studentEnrollments)).Msgf("Successfully retrieved %d student classes for student %d", len(studentEnrollments), studentID)

	classDataIDs := make(map[int]bool)

	for _, item := range studentEnrollments {
		classDataIDs[item.ClassID] = true
	}

	classIDs := pkg.GetKeysFromMap(classDataIDs)

	classes, err := s.practicumClassService.GetClassByIDs(classIDs)
	if err != nil {
		return nil, err
	}

	classMap := make(map[int]model.PracticumClass)
	for _, item := range classes {
		classMap[item.IDPracticumClass] = item
	}

	studentSchedules := make([]dto.StudentSchedules, len(studentEnrollments))
	for i, item := range studentEnrollments {
		studentClassObj := dto.StudentSchedules{
			ID:        item.ID,
			ClassName: classMap[item.ClassID].Name,
			ClassTime: classMap[item.ClassID].Time,
			Day:       classMap[item.ClassID].Day,
		}
		studentSchedules[i] = studentClassObj
	}

	return studentSchedules, nil

}
