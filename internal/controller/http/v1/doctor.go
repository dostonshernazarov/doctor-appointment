package v1

import (
	"time"

	"github.com/dostonshernazarov/doctor-appointment/internal/controller/http/models"
	"github.com/dostonshernazarov/doctor-appointment/internal/entity"
	"github.com/gofiber/fiber/v2"
)

// @Summary Create doctor
// @Description Create doctor
// @Accept json
// @Produce json
// @Tags doctor
// @Param doctor body models.Doctor true "Doctor"
// @Success 201 {object} models.DoctorResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /doctors [post]
func (h *HandlerV1) CreateDoctor(c *fiber.Ctx) error {
	doctor := models.Doctor{}
	if err := c.BodyParser(&doctor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	timeNow := time.Now()

	err := h.Doctor.CreateDoctor(c.Context(), entity.Doctor{
		Name:           doctor.Name,
		Specialization: doctor.Specialization,
		Schedule:       doctor.Schedule,
		CreatedAt:      timeNow,
		UpdatedAt:      timeNow,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(models.DoctorResponse{
		ID:             doctor.ID,
		Name:           doctor.Name,
		Specialization: doctor.Specialization,
		Schedule:       doctor.Schedule,
		CreatedAt:      timeNow,
		UpdatedAt:      timeNow,
	})
}

// @Summary Get doctor by id
// @Description Get doctor by id
// @Accept json
// @Produce json
// @Tags doctor
// @Param id path int true "Doctor ID"
// @Success 200 {object} models.DoctorResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /doctors/{id} [get]
func (h *HandlerV1) GetDoctorByID(c *fiber.Ctx) error {
	doctorID := c.Params("id")

	doctor, err := h.Doctor.GetDoctorByID(c.Context(), doctorID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(models.DoctorResponse{
		ID:             doctor.ID,
		Name:           doctor.Name,
		Specialization: doctor.Specialization,
		Schedule:       doctor.Schedule,
		CreatedAt:      doctor.CreatedAt,
		UpdatedAt:      doctor.UpdatedAt,
	})
}

// @Summary Update doctor
// @Description Update doctor
// @Accept json
// @Produce json
// @Tags doctor
// @Param id path int true "Doctor ID"
// @Param doctor body models.Doctor true "Doctor"
// @Success 200 {object} models.Doctor
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /doctors/{id} [put]
func (h *HandlerV1) UpdateDoctor(c *fiber.Ctx) error {
	doctorID := c.Params("id")

	doctor := models.Doctor{}
	if err := c.BodyParser(&doctor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Get doctor by id
	doctorGet, err := h.Doctor.GetDoctorByID(c.Context(), doctorID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	timeNow := time.Now()

	err = h.Doctor.UpdateDoctor(c.Context(), entity.Doctor{
		ID:             doctorID,
		Name:           doctor.Name,
		Specialization: doctor.Specialization,
		Schedule:       doctor.Schedule,
		CreatedAt:      doctorGet.CreatedAt,
		UpdatedAt:      timeNow,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(models.DoctorResponse{
		ID:             doctorID,
		Name:           doctor.Name,
		Specialization: doctor.Specialization,
		Schedule:       doctor.Schedule,
		CreatedAt:      doctorGet.CreatedAt,
		UpdatedAt:      timeNow,
	})
}

// @Summary Delete doctor
// @Description Delete doctor
// @Accept json
// @Produce json
// @Tags doctor
// @Param id path int true "Doctor ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /doctors/{id} [delete]
func (h *HandlerV1) DeleteDoctor(c *fiber.Ctx) error {
	doctorID := c.Params("id")

	err := h.Doctor.DeleteDoctor(c.Context(), doctorID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(models.SuccessResponse{
		Message: "Doctor deleted successfully",
	})
}

// @Summary Get all doctors
// @Description Get all doctors
// @Accept json
// @Produce json
// @Tags doctor
// @Success 200 {object} models.AllDoctorsResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /doctors [get]
func (h *HandlerV1) GetAllDoctors(c *fiber.Ctx) error {
	doctors, err := h.Doctor.GetAllDoctors(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(models.AllDoctorsResponse{
		Doctors: doctors,
	})
}

// @Summary Get doctors by specialization
// @Description Get doctors by specialization
// @Accept json
// @Produce json
// @Tags doctor
// @Param specialization path string true "Specialization"
// @Success 200 {object} models.AllDoctorsResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /doctors/specialization/{specialization} [get]
func (h *HandlerV1) GetDoctorsBySpecialization(c *fiber.Ctx) error {
	specialization := c.Params("specialization")

	doctors, err := h.Doctor.GetDoctorsBySpecialization(c.Context(), specialization)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(models.AllDoctorsResponse{
		Doctors: doctors,
	})
}

// @Summary List specializations
// @Description List specializations
// @Accept json
// @Produce json
// @Tags doctor
// @Success 200 {object} models.SpecializationResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /doctors/specializations [get]
func (h *HandlerV1) ListSpecializations(c *fiber.Ctx) error {
	specializations, err := h.Doctor.ListSpecializations(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(models.SpecializationResponse{
		Specializations: specializations,
	})
}

// @Summary Get booked schedules by doctor id
// @Description Get booked schedules by doctor id
// @Accept json
// @Produce json
// @Tags doctor
// @Param id path int true "Doctor ID"
// @Success 200 {object} models.BookedSchedulesResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /doctors/{id}/booked-schedules [get]
func (h *HandlerV1) GetBookedSchedulesByDoctorID(c *fiber.Ctx) error {
	doctorID := c.Params("id")

	bookedSchedules, err := h.Doctor.GetBookedSchedulesByDoctorID(c.Context(), doctorID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(models.BookedSchedulesResponse{
		BookedSchedules: bookedSchedules,
	})
}
