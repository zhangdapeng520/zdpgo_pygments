from pathlib import Path

BASE_DIR = Path(__file__).resolve().parent.parent

SECRET_KEY = "django-insecure-)*#p7$3!ke92gc7-_ufvm#pr&g-b5@qee#lfu-rl%8ecgnv=tm"
DEBUG = True

ALLOWED_HOSTS = []

MIDDLEWARE = [
    "django.middleware.security.SecurityMiddleware",
    "django.contrib.sessions.middleware.SessionMiddleware",
    "corsheaders.middleware.CorsMiddleware",  # 跨域请求中间件
    "django.middleware.common.CommonMiddleware",
    "django.middleware.csrf.CsrfViewMiddleware",
    "django.contrib.auth.middleware.AuthenticationMiddleware",
    "django.contrib.messages.middleware.MessageMiddleware",
    "django.middleware.clickjacking.XFrameOptionsMiddleware",
]

# 应用列表
INSTALLED_APPS = [
    "django.contrib.admin",
    "django.contrib.auth",
    "django.contrib.contenttypes",
    "django.contrib.sessions",
    "django.contrib.messages",
    "django.contrib.staticfiles",
    "django.contrib.sites",  # 站点设置，用于支持django-allauth

    # 第三方应用
    "rest_framework",
    "corsheaders",
    "rest_framework.authtoken",  # 权限token
    "allauth",  # django-allauth权限
    "allauth.account",  # 权限账号
    "allauth.socialaccount",  # 社交账号
    "dj_rest_auth",
    "dj_rest_auth.registration",  # 辅助实现注册

    # 本地应用
    "accounts.apps.AccountsConfig",
    "posts.apps.PostsConfig",
]

ROOT_URLCONF = "django_project.urls"

# 模板相关配置
# 模板相关配置
# 模板相关配置
TEMPLATES = [
    {
        "BACKEND": "django.template.backends.django.DjangoTemplates",
        "DIRS": [],
        "APP_DIRS": True,
        "OPTIONS": {
            "context_processors": [
                "django.template.context_processors.debug",
                "django.template.context_processors.request",
                "django.contrib.auth.context_processors.auth",
                "django.contrib.messages.context_processors.messages",
                # 用于支持django-allauth
                "django.template.context_processors.request",
            ],
        },
    },
]

WSGI_APPLICATION = "django_project.wsgi.application"

# Database
# https://docs.djangoproject.com/en/4.0/ref/settings/#databases
# https://docs.djangoproject.com/en/4.0/ref/settings/#databases
# https://docs.djangoproject.com/en/4.0/ref/settings/#databases
# https://docs.djangoproject.com/en/4.0/ref/settings/#databases
# https://docs.djangoproject.com/en/4.0/ref/settings/#databases
# https://docs.djangoproject.com/en/4.0/ref/settings/#databases
# https://docs.djangoproject.com/en/4.0/ref/settings/#databases
# https://docs.djangoproject.com/en/4.0/ref/settings/#databases
# https://docs.djangoproject.com/en/4.0/ref/settings/#databases
# https://docs.djangoproject.com/en/4.0/ref/settings/#databases
# https://docs.djangoproject.com/en/4.0/ref/settings/#databases
# https://docs.djangoproject.com/en/4.0/ref/settings/#databases
# https://docs.djangoproject.com/en/4.0/ref/settings/#databases

DATABASES = {
    "default": {
        "ENGINE": "django.db.backends.sqlite3",
        "NAME": BASE_DIR / "db.sqlite3",
    }
}

# Password validation
# https://docs.djangoproject.com/en/4.0/ref/settings/#auth-password-validators

AUTH_PASSWORD_VALIDATORS = [
    {
        "NAME": "django.contrib.auth.password_validation.UserAttributeSimilarityValidator",
    },
    {
        "NAME": "django.contrib.auth.password_validation.MinimumLengthValidator",
    },
    {
        "NAME": "django.contrib.auth.password_validation.CommonPasswordValidator",
    },
    {
        "NAME": "django.contrib.auth.password_validation.NumericPasswordValidator",
    },
]

# 配置语言：https://docs.djangoproject.com/en/4.0/topics/i18n/
# 配置语言：https://docs.djangoproject.com/en/4.0/topics/i18n/
# 配置语言：https://docs.djangoproject.com/en/4.0/topics/i18n/
# 配置语言：https://docs.djangoproject.com/en/4.0/topics/i18n/
# 配置语言：https://docs.djangoproject.com/en/4.0/topics/i18n/
# 配置语言：https://docs.djangoproject.com/en/4.0/topics/i18n/
# 配置语言：https://docs.djangoproject.com/en/4.0/topics/i18n/
# 配置语言：https://docs.djangoproject.com/en/4.0/topics/i18n/
# 配置语言：https://docs.djangoproject.com/en/4.0/topics/i18n/
# 配置语言：https://docs.djangoproject.com/en/4.0/topics/i18n/
# 配置语言：https://docs.djangoproject.com/en/4.0/topics/i18n/
# 配置语言：https://docs.djangoproject.com/en/4.0/topics/i18n/
# 配置语言：https://docs.djangoproject.com/en/4.0/topics/i18n/
# 配置语言：https://docs.djangoproject.com/en/4.0/topics/i18n/
# 配置语言：https://docs.djangoproject.com/en/4.0/topics/i18n/
# 配置语言：https://docs.djangoproject.com/en/4.0/topics/i18n/
# 配置语言：https://docs.djangoproject.com/en/4.0/topics/i18n/
LANGUAGE_CODE = 'zh-hans'
TIME_ZONE = 'Asia/Shanghai'
USE_I18N = True
USE_TZ = True

# Static files (CSS, JavaScript, Images)
# https://docs.djangoproject.com/en/4.0/howto/static-files/
STATIC_URL = "static/"

# Default primary key field type
# https://docs.djangoproject.com/en/4.0/ref/settings/#default-auto-field

DEFAULT_AUTO_FIELD = "django.db.models.BigAutoField"

# 用户权限模型
AUTH_USER_MODEL = "accounts.CustomUser"

# DRF相关配置
REST_FRAMEWORK = {
    "DEFAULT_PERMISSION_CLASSES": [
        "rest_framework.permissions.IsAuthenticated",  # 用户权限校验
    ],
    "DEFAULT_AUTHENTICATION_CLASSES": [  # 权限类
        "rest_framework.authentication.SessionAuthentication",
        "rest_framework.authentication.TokenAuthentication",  # new
    ],
}

# 跨域站点
# 跨域站点
# 跨域站点
# 跨域站点
# 跨域站点
# 跨域站点
# 跨域站点
# 跨域站点
# 跨域站点
# 跨域站点
# 跨域站点
# 跨域站点
# 跨域站点
# 跨域站点
# 跨域站点
CORS_ORIGIN_WHITELIST = (
    "http://localhost:3000",
    "http://localhost:8000",
)

# CSRF信任站点
CSRF_TRUSTED_ORIGINS = ["http://localhost:3000"]

# 邮件后端
EMAIL_BACKEND = "django.core.mail.backends.console.EmailBackend"  # new

# 站点ID
SITE_ID = 1
