package controller

import (
	"net/http"

	"github.com/B6025212/team05/entity"
	"github.com/asaskevich/govalidator"

	"github.com/gin-gonic/gin"
)



type extendedAdding_point struct {
	entity.Adding_point
	Professor_Name   string
	Student_ID 		 string
	Student_Name	 string
	Subject_ID		 string
	Section			 uint	
	Course_Name     string
	Subject_TH_Name string
	Subject_EN_Name string
	Day             string
	Start_Time      string
	End_Time        string
	Exam_Date       string
	Unit            string
	Exam_Start_Time string
	Exam_End_Time   string
}

// POST /course
func CreateAdding_point(c *gin.Context) {
	var adding_point entity.Adding_point
	var enroll entity.Enroll
	var professor entity.Professor
	var grade entity.Grade

	if err := c.ShouldBindJSON(&adding_point); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Communication Diagram Step
	// ค้นหา entity request ด้วย id ของ request ที่รับเข้ามา
	// SELECT * FROM requests` WHERE request_id = <request.Request_ID>
	// if tx := entity.DB().Where("adding_point_id = ?", adding_point.Adding_point_ID).First(&adding_point); tx.RowsAffected == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Adding_point not found"})
	// 	return
	// }
	if tx := entity.DB().Where("Professor_ID = ?", adding_point.Professor_ID).First(&professor); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "professor type not found"})
		return
	}
	// Communication Diagram Step
	// ค้นหา entity student ด้วย id ของ student ที่รับเข้ามา
	// SELECT * FROM `student` WHERE student_id = <student.Student_ID>
	if tx := entity.DB().Where("enroll_id = ?", adding_point.Enroll_ID).First(&enroll); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "enroll status not found"})
		return
	}

	// Communication Diagram Step
	// ค้นหา entity subject ด้วย id ของ subject ที่รับเข้ามา
	// SELECT * FROM `subject` WHERE subject_id = <subject.subject_id>
	

	// Communication Diagram Step
	// ค้นหา entity request_type ด้วย id ของ request_type ที่รับเข้ามา
	// SELECT * FROM `request_type` WHERE request_type_id = <request_type.request_type_ID>
	if tx := entity.DB().Where("grade_id = ?", adding_point.Grade_ID).First(&grade); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "grade not found"})
		return
	}


	if _, err := govalidator.ValidateStruct(adding_point); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	new_adding_point := entity.Adding_point{
		Adding_point_ID:  adding_point.Adding_point_ID,
		Professor: professor,
		Enroll_ID: adding_point.Enroll_ID,
		Grade_ID: adding_point.Grade_ID,
		
	}

	// บันทึก entity request
	if err := entity.DB().Create(&new_adding_point).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": new_adding_point})
}
func GetSubjectByAdding(c *gin.Context) {
	/* Query subject record(s) by subject_id */

	var extendedAdding_point []extendedAdding_point

	subject_id := c.Param("subject_id")

	query := entity.DB().Raw("SELECT e.*, s.* ,cs.day,cs.start_time,cs.end_time ,ex.* ,p.* FROM enrolls e JOIN subjects s JOIN class_schedules cs JOIN exam_schedules ex JOIN professors p ON e.subject_id = s.subject_id AND e.section = s.section AND e.class_schedule_id = cs.class_schedule_id AND e.exam_schedule_id = ex.exam_schedule_id AND s.professor_id = p.id WHERE subject_id = ?", subject_id).Scan(&extendedAdding_point)
	if err := query.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": extendedAdding_point})
}
func ListAddingByEnroll(c *gin.Context) {
	/*	++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
		Function for list all records from `subject` table.
		HTTP GET : /subjects
		++++++++++++++++++++++++++++++++++++++++++++++++++++++++++	*/

	//var enroll []entity.Enroll

	//SELECT e.*, c.* ,cs.* FROM `enrolls` e JOIN `subjects` c  JOIN `class_schedules` cs ON e.subject_id = c.subject_id  AND  e.section = c.section AND e.subject_id = cs.subject_id
	//SELECT e.* , c.* , s.* , ex.* , sj.* FROM enrolls e JOIN class_schedules c JOIN students s JOIN subjects sj JOIN exam_schedules ex ON e.section = sj.section AND e.student_id = s.student_id AND e.subject_id = sj.subject_id AND c.class_schedule_id = e.class_schedule_id AND s.student_id = e.student_id AND ex.subject_id = c.subject_id
	//SELECT e.* , c.* , s.* , ex.* , sj.* FROM enrolls e JOIN class_schedules c JOIN students s JOIN subjects sj JOIN exam_schedules ex ON e.section = sj.section AND e.student_id = s.student_id AND e.subject_id = sj.subject_id AND c.class_schedule_id = e.class_schedule_id AND s.student_id = e.student_id AND ex.subject_id = c.subject_id
	//SELECT e.*, s.* ,cs.day,cs.start_time,cs.end_time FROM enrolls e JOIN subjects s JOIN class_schedules cs ON e.subject_id = s.subject_id AND e.section = s.section AND e.class_schedule_id = cs.class_schedule_id;
	var extendedAdding_point []extendedAdding_point
	query := entity.DB().Raw("SELECT e.*, s.* ,cs.day,cs.start_time,cs.end_time ,ex.* ,p.* FROM enrolls e JOIN subjects s JOIN class_schedules cs JOIN exam_schedules ex JOIN professors p ON e.subject_id = s.subject_id AND e.section = s.section AND e.class_schedule_id = cs.class_schedule_id AND e.exam_schedule_id = ex.exam_schedule_id AND s.professor_id = p.id").Scan(&extendedAdding_point)
	if err := query.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": extendedAdding_point})

}

// List /request
func ListAdding_point(c *gin.Context) {
	var  extendedAdding_point []extendedAdding_point
	if err := entity.DB().Raw("SELECT a.*, s.*, at.*,c.*,sd.* FROM adding_points a JOIN grades at JOIN enrolls c JOIN students sd JOIN subjects s ON a.enroll_id = c.enroll_id AND  sd.student_id = c.student_id AND c.subject_id = s.subject_id AND   s.section = c.section AND   a.grade_id = at.grade_id  ").Scan(&extendedAdding_point).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": extendedAdding_point})
}

// Get /request
func GetAdding_point(c *gin.Context) {
	var adding_point entity.Adding_point
	id := c.Param("adding_point_id")
	if err := entity.DB().Where("adding_point_id = ?", id).First(&adding_point).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": adding_point})
}

// // DELETE /request
func DeleteAdding_point(c *gin.Context) {

	id := c.Param("Adding_point_ID")

	if tx := entity.DB().Exec("DELETE FROM adding_points WHERE Adding_point_ID = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /professors
func UpdateAdding_point(c *gin.Context) {
	var adding_point entity.Adding_point
	var enroll entity.Enroll
	var professor entity.Professor
	var grade entity.Grade

	if err := c.ShouldBindJSON(&adding_point); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("enroll_id = ?", adding_point.Enroll_ID).First(&enroll); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "enroll not found"})
		return
	}
	if tx := entity.DB().Where("professor_id = ?", adding_point.Professor_ID).First(&professor); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "professor not found"})
		return
	}
	// if tx := entity.DB().Where("subject_id = ? AND section = ?", request.Subject_ID, subject.Section).Find(&subject); tx.RowsAffected == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "subject with this section not found"})
	// 	return
	// }
	if tx := entity.DB().Where("grade_id = ?", adding_point.Grade_ID).First(&grade); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "grade type not found"})
		return
	}

		updated_adding_point := entity.Adding_point{
		Adding_point_ID:  adding_point.Adding_point_ID,
		Grade: adding_point.Grade,
		Enroll: enroll,
		Professor: professor,
		
	}

	
	if err := entity.DB().Save(&updated_adding_point).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": updated_adding_point})

}
func GetPreviousAdding_point(c *gin.Context) {
	var adding_point entity.Adding_point
	if err := entity.DB().Last(&adding_point).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": adding_point})
}

// //รับค่าprofessorมากรองรหัสวิชาและกลุ่ม
func GetStudenByEnroll(c *gin.Context) {
	var  enroll []entity.Enroll
	section := c.Param("section")
	subject_id := c.Param("subject_id")
	query := entity.DB().Where("subject_id = ? AND section = ?",subject_id,section).First(&enroll)
	if err := query.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": enroll})
	// if err := entity.DB().Raw("SELECT g.*, a.*,e.*,s.*,sn.* FROM adding_points a  JOIN grades g  JOIN enrolls e JOIN subjects s JOIN students sn ON g.grade_id = a.grade_id AND a.enroll_id = e.enroll_id AND e.subject_id = s.subject_id AND e.section = s.section AND e.student_id =sn.student_id", subject_id,section).Find(&enroll).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{"data": enroll})
}




// //เอาค่ากลุ่มกับรหัสมากรองชื่อทั้งหมด
// func GetSubjectByProfessor(c *gin.Context) {
// 	var  subject []entity.Subject
// 	professor_id := c.Param("professor_id")
// 	//section := c.Param("section")
// 	if err := entity.DB().Raw(" SELECT * FROM subjects where professor_id =?", professor_id ).Find(&subject).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": subject})
// }
