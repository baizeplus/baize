package systemServiceImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemDao/systemDaoImpl"
	"baize/app/business/system/systemModels"
	"baize/app/utils/snowflake"
	"github.com/baizeplus/sqly"
	"github.com/gin-gonic/gin"
	"strconv"
)

type DeptService struct {
	data    *sqly.DB
	deptDao systemDao.IDeptDao
	roleDao systemDao.IRoleDao
}

func NewDeptService(data *sqly.DB, dd *systemDaoImpl.SysDeptDao, rd *systemDaoImpl.SysRoleDao) *DeptService {
	return &DeptService{data: data, deptDao: dd, roleDao: rd}
}

func (ds *DeptService) SelectDeptList(c *gin.Context, dept *systemModels.SysDeptDQL) (list []*systemModels.SysDeptVo) {
	return ds.deptDao.SelectDeptList(c, ds.data, dept)

}

func (ds *DeptService) SelectDeptById(c *gin.Context, deptId int64) (dept *systemModels.SysDeptVo) {
	return ds.deptDao.SelectDeptById(c, ds.data, deptId)

}

func (ds *DeptService) InsertDept(c *gin.Context, dept *systemModels.SysDeptVo) {
	//获取上级部门祖籍信息
	parentDept := ds.SelectDeptById(c, dept.ParentId)
	dept.Ancestors = parentDept.Ancestors + "," + strconv.FormatInt(dept.ParentId, 10)
	//执行添加
	dept.DeptId = snowflake.GenID()
	ds.deptDao.InsertDept(c, ds.data, dept)
	return
}

func (ds *DeptService) UpdateDept(c *gin.Context, dept *systemModels.SysDeptVo) {
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
