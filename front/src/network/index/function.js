import header from "./header";

export default {
    insertDiff, insertProblem,
    deleteProblem,
    updataProblem,
    queryAllType, queryAllDiffical, queryAllProblem
}

// 添加题目难度 {"diff" : (int)}
export function insertDiff( self, data ) {
    header.insertDifficulty(data).then((res) => {
        res = res.data
        if(res.data == 0) {
            self.$message.err("题目难度添加失败..."); return
        }
        self.$message.success("题目难度添加成功 :)")
    }).catch((err) => {
        self.$message.err("服务器异常" + err);
    });
}
// 添加题目 {"data" : { Problem_name : "", Problem_url : "", Type_id : (int), Solve_time : ""(yyyy-mm-dd), Diff_id : (int) }}
export function insertProblem( self, data ) {
    header.insertProblem(data).then((res) => {
        res = res.data;
        if(res.code == 0) {
            self.$message.error("问题添加失败..."); return;
        }
        self.$message.success("问题添加成功 :)")
    }).catch((err) => {
        self.$message.err("服务器异常" + err);
    });    
}

// 删除问题 {"Problem_id" : (int)}
export function deleteProblem( self, data ) {
    header.deleteProblem(data).then((res) => {
        res = res.data;
        if(res.code == 0) {
            self.$message.error("问题删除失败..."); return;
        }
        self.$message.success("问题删除成功 :)")
    }).catch((err) => {
        self.$message.err("服务器异常" + err);
    }); 
}

// 修改问题 {"data" : { Problem_name : "", Problem_url : "", Type_id : (int), Solve_time : ""(yyyy-mm-dd), Diff_id : (int) }, "Problem_id" : (int)}
export function updataProblem( self, data ) {
    header.updataProblem(data).then((res) => {
        res = res.data;
        if(res.code == 0) {
            self.$message.error("问题修改失败..."); return;
        }
        self.$message.success("问题修改成功 :)")
    }).catch((err) => {
        self.$message.err("服务器异常" + err);
    });
}

// 查询所有的题目来源
export function queryAllType( self ) {
    return header.queryAllType().then((res) => {
      res = res.data
      if(res.code == 0) {
        self.$message.err("读取全部题目来源失败..."); return null
      }
      return res.data
    }).catch((err) => {
      self.$message.err("服务器异常" + err);
    });
}
// 查询所有的题目困难度
export function queryAllDiffical( self ) {
    return header.queryAllDiffical().then((res) => {
      res = res.data
      if(res.code == 0) {
        self.$message.err("读取全部题目难度失败..."); return null
      }
      return res.data
    }).catch((err) => {
        self.$message.err("服务器异常" + err);
    });
}
// 查询所有的题目
export function queryAllProblem( self ) {
    return header.queryAllProblem().then(res => {
        res = res.data
        if(res.code == 0) {
            self.$message.err("读取全部题目失败..."); return null
          }
        return res.data
    }).catch((err) => {
        self.$message.err("服务器异常" + err);
    });
}