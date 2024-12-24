<script setup>
import { ref } from 'vue';

// 团队成员数据
const teamMembers = ref([
    { name: '韩梦群', position: '项目经理', intro: '我负责整个项目的规划和管理。' },
    { name: '刘瑞琪', position: '前端开发', intro: '我负责网站的前端开发和界面设计。' },
    { name: '林星宇', position: '后端开发', intro: '我负责服务器端逻辑和数据库设计。' },
    { name: '吴昊阳', position: 'UI设计师', intro: '我负责平台的UI设计和用户体验优化。' },
    { name: '赵孜阳', position: '运行维护', intro: '我负责平台的运行维护和部署。' },
]);

</script>

<template>

    <div>
        <header class="about-header">
            <div class="container m-top-20">
                <h1 class="logo">
                    <RouterLink to="/">淘山大</RouterLink>
                </h1>
                <RouterLink class="entry" to="/">
                    进入网站首页
                    <i class="iconfont icon-angle-right"></i>
                    <i class="iconfont icon-angle-right"></i>
                </RouterLink>
            </div>
        </header>

        <section class="about-us-page">
            <header class="header">
                <h1>关于我们</h1>
            </header>

            <section class="content">
                <div class="section">
                    <h2>一、项目简介</h2>
                    <p>
                        本平台是为了为山东大学的学生提供一个便捷的二手交易平台，支持学弟学妹们快速发布和购买二手商品。
                        我们通过平台实现信息共享，促进校园内物品流通，提升资源利用效率，帮助学生节省开销。
                    </p>
                </div>

                <div class="section">
                    <h2>二、团队介绍</h2>
                    <el-table :data="teamMembers" style="width: 100%">
                        <el-table-column prop="name" label="姓名" width="150"></el-table-column>
                        <el-table-column prop="position" label="职位" width="150"></el-table-column>
                        <el-table-column prop="intro" label="自我介绍"></el-table-column>
                    </el-table>
                </div>
            </section>
        </section>
        <footer class="login-footer">
            <div class="container">
                <p>
                    <a href="javascript:;">关于我们</a>
                    <a href="javascript:;">帮助中心</a>
                    <a href="javascript:;">售后服务</a>
                    <a href="javascript:;">商务合作</a>
                    <a href="javascript:;">搜索推荐</a>
                    <a href="javascript:;">友情链接</a>
                </p>
                <p>CopyRight &copy; 淘山大</p>
            </div>
        </footer>
    </div>


</template>

<style scoped lang="scss">
.about-header {
    background: #fff;
    border-bottom: 1px solid #e4e4e4;

    .container {
        display: flex;
        align-items: center;
        justify-content: space-between;
    }

    .logo {
        width: 600px;

        a {
            display: block;
            height: 130px;
            width: 100%;
            text-indent: -9999px;
            background: url("@/assets/images/logo11.png") no-repeat center 30px / contain;
            background-size: 360px auto;
        }
    }

    .sub {
        flex: 1;
        font-size: 24px;
        font-weight: normal;
        margin-bottom: 38px;
        margin-left: 20px;
        color: #666;
    }

    .entry {
        width: 330px;
        margin-bottom: 50px;
        font-size: 18px;

        i {
            font-size: 14px;
            color: $mainColor;
            letter-spacing: -5px;
        }
    }
}

.about-us-page {
    background: url('@/assets/images/login-bg.png') no-repeat center / cover;
    height: 690px;
    position: relative;
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
    transition: all 0.3s ease;
}

.header {
    text-align: center;
    margin-bottom: 20px;
    padding-top: 20px;

    h1 {
        font-size: 26px;
        font-weight: bold;
        color: #333;
    }
}

.content {
    background: #fff;
    padding: 30px;
    border-radius: 12px;
    width: 100%;
    max-width: 1000px;
    margin: 0 auto
}

.content:hover {
    box-shadow: 0 15px 40px rgba(0, 0, 0, 0.15);
}

.section {
    margin-bottom: 40px;

    h2 {
        font-size: 20px;
        font-weight: 600;
        color: #409EFF;
        margin-bottom: 20px;
        border-bottom: 2px solid #409EFF;
        padding-bottom: 10px;
    }

    p {
        font-size: 18px;
        line-height: 1.8;
        color: #666;
    }
}

.el-table {
    margin-top: 20px;
    border-radius: 8px;
    overflow: hidden;
}

.el-table-column {
    text-align: center;
    font-size: 16px;
}

.el-table .cell {
    text-align: center;
    padding: 15px;
}

.el-table thead {
    background-color: #f2f2f2;
    font-weight: bold;
    color: #333;
}

.el-table .el-table-column .cell {
    border-top: 1px solid #f2f2f2;
    border-right: 1px solid #f2f2f2;
}

.el-table .el-table-column:last-child .cell {
    border-right: none;
}

.el-table .el-table-column .cell:last-child {
    border-right: none;
}

.login-footer {
    padding: 30px 0 50px;
    background: #fff;

    p {
        text-align: center;
        color: #999;
        padding-top: 20px;

        a {
            line-height: 1;
            padding: 0 10px;
            color: #999;
            display: inline-block;

            ~a {
                border-left: 1px solid #ccc;
            }
        }
    }
}


</style>
