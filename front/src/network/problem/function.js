import header from './header';

export default {
    changeSolution, changeDescription, 
    queryMdDetail,
    problemCrawler
}

// 修改题解
export function changeSolution( data, self ) {
    header.changeSolution(data).then((res)=> {
        res = res.data;
        if(res.code == 0) {
            self.$message.error("保存失败..."); return
        }
        self.$message.success("保存成功 :)");
    }).catch((err) => {
        self.$message.error("服务器异常 " + err);
    }); 
}
// 修改题目描述
export function changeDescription( self, data ) {
    header.changeDescription(data).then((res)=> {
        res = res.data;
        if(res.code == 0) {
            self.$message.error("保存失败..."); return
        }
        self.$message.success("保存成功 :)");
    }).catch((err) => {
        self.$message.error("服务器异常 " + err);
    }); 
}

// 查询问题的题目描述和题解 {"Problem_id": id}
export function queryMdDetail( self, data ) {
    return header.queryMdDetail(data).then((res) => {
        res = res.data;
        if(res.data == 0) {
            self.$message.error("题目详情获取失败..."); return
        }
        return res.data
    }).catch((err) => {
        self.$message.error("服务器异常 " + err);
    }); 
}

// 题目爬虫 {"Type" : "", "Problem_url" : ""}
export function problemCrawler( self, data ) {
    return header.problemCrawler(data).then((res) => {
        res = res.data
        if(res.code == 0) {
            self.$message.error("爬取失败，该类型爬虫不存在"); return
        }
        self.$message.success("爬取成功")
        return res.data
    }).catch((err) => {
        self.$message.error("服务器异常 " + err);
    }); 
}