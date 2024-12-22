<script setup></script>

<template>
    <div class="container">
        <div class="user-aside">
            <div class="user-manage">
                <h4>我的账户</h4>
                <div class="links">
                    <RouterLink to="/user">个人中心</RouterLink>
                </div>
                <h4>交易管理</h4>
                <div class="links">
                    <RouterLink to="/user/buyer_order">购入订单</RouterLink>
                    <RouterLink to="/user/seller_order">售出订单</RouterLink>
                    <RouterLink to="/user/good">我的商品</RouterLink>
                </div>
                <h4>售后管理</h4>
                <div class="links">
                    <RouterLink to="/user/problem">我的售后</RouterLink>
                </div>
            </div>
        </div>
        <div class="article">
            <!-- 三级路由的挂载点 -->
            <RouterView />
        </div>
    </div>
</template>

<style scoped lang="scss">
.container {
    display: flex;
    padding-top: 20px;

    .user-aside {
        width: 250px;
        margin-right: 20px;
        border-radius: 2px;
        background-color: #fff;

        .user-manage {
            background-color: #fff;

            h4 {
                font-size: 20px;
                font-weight: 400;
                padding: 20px 52px 5px;
                border-top: 1px solid #f6f6f6;
            }

            .links {
                padding: 0 52px 10px;
            }

            a {
                display: block;
                line-height: 1;
                padding: 15px 0;
                font-size: 16px;
                color: #666;
                position: relative;

                &:hover {
                    color: $mainColor;
                }

                &.active,
                &.router-link-exact-active {
                    color: $mainColor;

                    &:before {
                        display: block;
                    }
                }

                &:before {
                    content: '';
                    display: none;
                    width: 6px;
                    height: 6px;
                    border-radius: 50%;
                    position: absolute;
                    top: 19px;
                    left: -16px;
                    background-color: $mainColor;
                }
            }
        }
    }

    .article {
        width: 1300px;
        background-color: #fff;
    }
}
</style>