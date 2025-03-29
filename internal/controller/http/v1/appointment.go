package v1

import (
	"github.com/gofiber/fiber/v2"
)

// @Summary Create appointment
// @Description Create appointment
// @Accept json
// @Produce json
// @Tags appointment
// @Param appointment body models.Appointment true "Appointment"
// @Success 201 {object} models.AppointmentResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /appointments [post]
func (h *HandlerV1) CreateAppointment(c *fiber.Ctx) error {
	appointment := models.Appointment{}
	if err := c.BodyParser(&appointment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	appointment.Status = "scheduled"

	err := h.Appointment.CreateAppointment(c.Context(), entity.Appointment{
		DoctorID:        appointment.DoctorID,
		UserID:          appointment.UserID,
		AppointmentTime: appointment.AppointmentTime,
		Duration:        appointment.Duration,
		Status:          appointment.Status,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(models.AppointmentResponse{
		ID:       appointment.ID,
		DoctorID: appointment.DoctorID,
		UserID:   appointment.UserID,
	})
}

// @Summary Get appointments by doctor id
// @Description Get appointments by doctor id
// @Accept json
// @Produce json
// @Tags appointment
// @Param doctor_id path int true "Doctor ID"
// @Success 200 {object} models.AppointmentResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /appointments/doctor/{doctor_id} [get]
func (h *HandlerV1) GetAppointmentsByDoctorID(c *fiber.Ctx) error {
	doctorID := c.Params("doctor_id")

	appointments, err := h.Appointment.GetAppointmentsByDoctorID(c.Context(), doctorID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(models.AppointmentsResponse{
		Appointments: appointments,
	})
}

// @Summary Get appointments by user id
// @Description Get appointments by user id
// @Accept json
// @Produce json
// @Tags appointment
// @Param user_id path int true "User ID"
// @Success 200 {object} models.AppointmentResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /appointments/user/{user_id} [get]
func (h *HandlerV1) GetAppointmentsByUserID(c *fiber.Ctx) error {
	userID := c.Params("user_id")

	appointments, err := h.Appointment.GetAppointmentsByUserID(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(models.AppointmentsResponse{
		Appointments: appointments,
	})
}

// @Summary Update appointment
// @Description Update appointment
// @Accept json
// @Produce json
// @Tags appointment
// @Param appointment body models.Appointment true "Appointment"
// @Success 200 {object} models.AppointmentResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /appointments [put]
func (h *HandlerV1) UpdateAppointment(c *fiber.Ctx) error {
	appointment := models.Appointment{}
	if err := c.BodyParser(&appointment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.Appointment.UpdateAppointment(c.Context(), entity.Appointment{
		ID:              appointment.ID,
		DoctorID:        appointment.DoctorID,
		UserID:          appointment.UserID,
		AppointmentTime: appointment.AppointmentTime,
		Duration:        appointment.Duration,
		Status:          appointment.Status,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(models.AppointmentResponse{
		ID:              appointment.ID,
		DoctorID:        appointment.DoctorID,
		UserID:          appointment.UserID,
		AppointmentTime: appointment.AppointmentTime,
		Duration:        appointment.Duration,
		Status:          appointment.Status,
	})
}

// @Summary Delete appointment
// @Description Delete appointment
// @Accept json
// @Produce json
// @Tags appointment
// @Param appointment_id path int true "Appointment ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /appointments/{appointment_id} [delete]
func (h *HandlerV1) DeleteAppointment(c *fiber.Ctx) error {
	appointmentID := c.Params("appointment_id")

	err := h.Appointment.DeleteAppointment(c.Context(), appointmentID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(models.SuccessResponse{
		Message: "Appointment deleted successfully",
	})
}

// @Summary Get booked schedules by doctor id
// @Description Get booked schedules by doctor id
// @Accept json
// @Produce json
// @Tags appointment
// @Param doctor_id path int true "Doctor ID"
// @Success 200 {object} models.AppointmentsResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /appointments/doctor/{doctor_id}/booked-schedules [get]
func (h *HandlerV1) GetBookedSchedulesByDoctorID(c *fiber.Ctx) error {
	doctorID := c.Params("doctor_id")

	bookedSchedules, err := h.Appointment.GetBookedSchedulesByDoctorID(c.Context(), doctorID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(models.AppointmentsResponse{
		Appointments: bookedSchedules,
	})
}

// @Summary Get booked schedules by user id
// @Description Get booked schedules by user id
// @Accept json
// @Produce json
// @Tags appointment
// @Param user_id path int true "User ID"
// @Success 200 {object} models.AppointmentsResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /appointments/user/{user_id}/booked-schedules [get]
func (h *HandlerV1) GetBookedSchedulesByUserID(c *fiber.Ctx) error {
	userID := c.Params("user_id")

	bookedSchedules, err := h.Appointment.GetBookedSchedulesByUserID(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(models.AppointmentsResponse{
		Appointments: bookedSchedules,
	})
}

// Get appointment by id
// @Summary Get appointment by id
// @Description Get appointment by id
// @Accept json
// @Produce json
// @Tags appointment
// @Param appointment_id path int true "Appointment ID"
// @Success 200 {object} models.AppointmentResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /appointments/{appointment_id} [get]
func (h *HandlerV1) GetAppointmentByID(c *fiber.Ctx) error {
	appointmentID := c.Params("appointment_id")

	appointment, err := h.Appointment.GetAppointmentByID(c.Context(), appointmentID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(models.AppointmentResponse{
		ID:              appointment.ID,
		DoctorID:        appointment.DoctorID,
		UserID:          appointment.UserID,
		AppointmentTime: appointment.AppointmentTime,
		Duration:        appointment.Duration,
		Status:          appointment.Status,
	})
}
