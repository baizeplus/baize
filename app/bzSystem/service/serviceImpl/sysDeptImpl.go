package serviceImpl

import (
	systemDao "baize/app/bzSystem/dao"
	"baize/app/bzSystem/dao/daoImpl"
	"baize/app/bzSystem/models"
	"baize/app/utils/snowflake"
	"github.com/baizeplus/sqly"
	"github.com/gin-gonic/gin"
)

type DeptService struct {
	data    *sqly.DB
	deptDao systemDao.IDeptDao
	roleDao systemDao.IRoleDao
}

func NewDeptService(data *sqly.DB, dd *daoImpl.SysDeptDao, rd *daoImpl.SysRoleDao) *DeptService {
	return &DeptService{data: data, deptDao: dd, roleDao: rd}
}

func (ds *DeptService) SelectDeptList(c *gin.Context, dept *models.SysDeptDQL) (list []*models.SysDeptVo) {
	return ds.deptDao.SelectDeptList(c, ds.data, dept)

}

func (ds *DeptService) SelectDeptById(c *gin.Context, deptId int64) (dept *models.SysDeptVo) {
	return ds.deptDao.SelectDeptById(c, ds.data, deptId)

}

func (ds *DeptService) InsertDept(c *gin.Context, dept *models.SysDeptVo) {
	dept.DeptId = snowflake.GenID()
	ds.deptDao.InsertDept(c, ds.data, dept)
	return
}

func (ds *DeptService) UpdateDept(c *gin.Context, dept *models.SysDeptVo) {
	ds.deptDao.UpdateDept(c, ds.data, dept)
}
func (ds *DeptService) DeleteDeptById(c *gin.Context, dept int64) {
	ds.deptDao.DeleteDeptById(c, ds.data, dept)
	return
}
func (ds *DeptService) CheckDeptNameUnique(c *gin.Context, id, parentId int64, deptName string) bool {
	deptId := ds.deptDao.CheckDeptNameUnique(c, ds.data, deptName, parentId)
	if deptId == id || deptId == 0 {
		return false
	}
	return true
}
func (ds *DeptService) HasChildByDeptId(c *gin.Context, deptId int64) bool {
	return ds.deptDao.HasChildByDeptId(c, ds.data, deptId) > 0
}

func (ds *DeptService) CheckDeptExistUser(c *gin.Context, deptId int64) bool {
	return ds.deptDao.CheckDeptExistUser(c, ds.data, deptId) > 0
}
func (ds *DeptService) SelectDeptListByRoleId(c *gin.Context, roleId int64) []string {
	role := ds.roleDao.SelectRoleById(c, ds.data, roleId)
	return ds.deptDao.SelectDeptListByRoleId(c, ds.data, roleId, role.DeptCheckStrictly)
}
