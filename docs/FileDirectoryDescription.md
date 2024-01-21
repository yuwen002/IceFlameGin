
    cmd 目录用于存放应用程序的入口文件。
    config 目录用于存放配置文件和配置相关的代码。
    docs 目录包含了项目的文档和设计文档。其中，api 目录用于存放 API 文档，design 目录用于存放项目的架构设计和数据流程图等文档。
    internal 目录是项目的主要实现目录，包含了业务逻辑的实现。
        app：存放应用程序的主要业务逻辑代码。
            controllers/：按照业务模块划分的控制器代码。
            middlewares 目录用于存放中间件代码。
            models 目录用于存放模型定义的代码。
            repositories 目录用于存放数据仓库的实现代码。
            services 目录用于存放服务层的实现代码。
        pkg：存放通用的、可复用的代码。
            database：数据库相关的代码。
            logger：日志记录器相关的代码。
            utils：通用的工具函数或帮助类。
    migrations 目录用于存放数据库迁移文件，包括升级和回滚脚本。
    routers 目录用于存放路由设置相关的代码。
    web 目录用于存放与 Web 相关的代码。
        handlers 目录用于存放处理请求的处理函数。
        static 目录用于存放静态资源文件，例如 CSS、JavaScript 等。
        templates 目录用于存放模板文件，例如 HTML 模板。

