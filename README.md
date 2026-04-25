# OpenCode Advanced 🚀

**تم تطويره بواسطة xbitoi** - نسخة متقدمة من OpenCode مع مميزات احترافية متقدمة

## 🎯 المميزات الرئيسية

### 1. 🤖 دعم متقدم للنماذج الذكية
- **Gemini Bridge**: جسر متقدم للاتصال بـ Google Gemini كمستخدم حقيقي
- **DeepSeek Bridge**: جسر للاتصال بـ DeepSeek مع تجاوز كشف الروبوتات
- **اختيار ديناميكي**: اختر بين النماذج المختلفة من قائمة سهلة الاستخدام
- **التوازي**: استخدم عدة نماذج في نفس الوقت

### 2. 💬 تكامل Telegram كامل
- تحكم كامل عبر Telegram Bot
- إرسال المشاريع والأسئلة مباشرة من Telegram
- تلقي النتائج والإشعارات على الفور
- دعم الملفات والصور والمستندات
- جلسات محادثة منفصلة لكل مستخدم

### 3. 🔒 نظام الأذونات المحسّن
- **الوضع الآمن**: فقط إذن واحد (الحذف)
- منع التنفيذ التلقائي للأوامر الخطيرة
- تصديق يدوي لكل عملية حساسة
- سجل شامل للعمليات المنفذة

### 4. 🧠 التعديل الذاتي والتعلم
- إضافة مهارات جديدة من خلال محادثة Telegram
- تحسين نفسي تلقائي للبرنامج
- حفظ وتحميل المهارات المخصصة
- دعم المكتبات والأدوات الخارجية
- نظام الذاكرة المشروع الذكية

### 5. 🏗️ أدوات بناء المشاريع المتقدمة
- **Project Builder**: بناء مشاريع معقدة من الصفر
- **Code Generator**: توليد كود احترافي
- **Architecture Assistant**: مساعد في تصميم البنية
- **Database Manager**: إدارة قوعد البيانات
- **DevOps Helper**: مساعد العمليات

### 6. 🖥️ التوافق الكامل مع Windows 11
- لا توجد أخطاء في التشغيل
- واجهة رسومية اختيارية
- دعم PowerShell و CMD
- تكامل سهل مع النظام
- دعم WSL و Windows Terminal

### 7. 🌐 مميزات إضافية احترافية
- **Project Memory**: ذاكرة المشروع الذكية
- **Multi-Session**: جلسات متعددة ومتوازية
- **Analytics**: تحليل الأداء والاستخدام
- **Plugin System**: نظام المكونات الإضافية
- **Version Control**: إدارة النسخ التلقائية
- **Code Review**: مراجعة الكود الذكية
- **Testing Framework**: إطار الاختبار المدمج
- **Documentation Auto-Generator**: توليد التوثيق التلقائي

## 📋 المتطلبات

- Go 1.24.0 أو أعلى
- Windows 11 / Linux / macOS
- Telegram Bot Token (للدعم عبر Telegram)
- API Keys (Gemini/DeepSeek اختياري)
- Chrome/Chromium مثبت (للبرايدج)

## 🚀 التثبيت السريع

```bash
# استنساخ المستودع
git clone https://github.com/xbitoi/opencode-advanced.git
cd opencode-advanced

# التثبيت
go install ./cmd/main.go

# أو البناء المباشر
go build -o opencode-advanced ./cmd/main.go

# التشغيل
./opencode-advanced
```

## ⚙️ التكوين

أنشئ ملف `.opencode.json` في المجلد الرئيسي:

```json
{
  "telegram": {
    "botToken": "YOUR_BOT_TOKEN_HERE",
    "adminChatID": "YOUR_CHAT_ID",
    "enabled": true
  },
  "bridges": {
    "gemini": {
      "enabled": true,
      "apiKey": "YOUR_GEMINI_API_KEY",
      "useBrowser": true,
      "headless": true
    },
    "deepseek": {
      "enabled": true,
      "browserMode": true,
      "antiDetect": true,
      "useProxy": false
    }
  },
  "permissions": {
    "allowDelete": true,
    "allowModify": false,
    "requireConfirmation": true,
    "adminOnly": false
  },
  "skills": {
    "directory": "./.opencode/skills",
    "autoLoad": true,
    "autoSave": true
  },
  "projects": {
    "directory": "./.opencode/projects",
    "autoBackup": true
  },
  "models": {
    "default": "gemini",
    "timeout": 300,
    "maxTokens": 8000
  },
  "windows": {
    "useGUI": false,
    "shell": "powershell"
  }
}
```

## 💬 الاستخدام عبر Telegram

```
/start                    - بدء المساعد
/ask <سؤال>               - طرح سؤال
/build <وصف>              - بناء مشروع جديد
/learn <مهارة>            - إضافة مهارة جديدة
/code <الكود>             - تحليل كود
/project <الأمر>          - إدارة المشاريع
/skills                   - عرض المهارات الحالية
/status                   - حالة النظام
/help                     - المساعدة الشاملة
/settings                 - الإعدادات
```

## 📟 استخدام سطر الأوامر

```bash
# وضع تفاعلي
opencode-advanced

# وضع غير تفاعلي
opencode-advanced -p "Your prompt here"

# تشغيل Telegram Bot فقط
opencode-advanced --telegram

# وضع التطوير مع التسجيل المفصل
opencode-advanced --debug

# اختيار نموذج محدد
opencode-advanced --model gemini

# إضافة مهارة
opencode-advanced --learn "skill_name" --describe "skill_description"
```

## 🏗️ البنية المعمارية

```
opencode-advanced/
├── cmd/
│   ├── main.go                 # نقطة الدخول الرئيسية
│   ├── telegram_handler.go     # معالج Telegram
│   └── flags.go                # معالج الأعلام
├── internal/
│   ├── bridges/
│   │   ├── gemini.go           # جسر Gemini
│   │   ├── deepseek.go         # جسر DeepSeek
│   │   ├── browser.go          # محرك المتصفح
│   │   └── antidetect.go       # تقنيات تجاوز الكشف
│   ├── telegram/
│   │   ├── client.go           # عميل Telegram
│   │   ├── commands.go         # معالجات الأوامر
│   │   ├── session.go          # إدارة الجلسات
│   │   └── notifications.go    # الإشعارات
│   ├── skills/
│   │   ├── manager.go          # مدير المهارات
│   │   ├── loader.go           # تحميل المهارات
│   │   ├── trainer.go          # تدريب المهارات
│   │   └── executor.go         # تنفيذ المهارات
│   ├── permissions/
│   │   ├── enforcer.go         # تطبيق الأذونات
│   │   ├── auditor.go          # تدقيق العمليات
│   │   └── rules.go            # قواعد الأذونات
│   ├── projects/
│   │   ├── builder.go          # بناء المشاريع
│   │   ├── generator.go        # توليد الكود
│   │   ├── manager.go          # إدارة المشاريع
│   │   └── templates.go        # قوالب المشاريع
│   ├── config/
│   │   ├── loader.go           # تحميل الإعدادات
│   │   ├── validator.go        # التحقق من الإعدادات
│   │   └── types.go            # أنواع البيانات
│   ├── models/
│   │   ├── selector.go         # اختيار النموذج
│   │   ├── manager.go          # إدارة النماذج
│   │   └── router.go           # توجيه الطلبات
│   ├── memory/
│   │   ├── project_memory.go   # ذاكرة المشروع
│   │   ├── session_memory.go   # ذاكرة الجلسة
│   │   └── storage.go          # التخزين الدائم
│   ├── analytics/
│   │   ├── tracker.go          # تتبع الاستخدام
│   │   ├── reporter.go         # تقارير الأداء
│   │   └── metrics.go          # المقاييس
│   ├── plugins/
│   │   ├── loader.go           # تحميل البرامج الإضافية
│   │   ├── registry.go         # تسجيل البرامج
│   │   └── executor.go         # تنفيذ البرامج
│   └── database/
│       ├── sqlite.go           # قاعدة البيانات
│       ├── migrations.go       # الهجرات
│       └── schema.go           # المخطط
├── plugins/                    # المكونات الإضافية الافتراضية
│   ├── code_reviewer.go
│   ├── doc_generator.go
│   └── test_runner.go
├── skills/                     # المهارات المخصصة
│   └── example_skill.go
├── config/
│   └── example.opencode.json   # ملف إعدادات مثال
├── tests/
│   ├── bridges_test.go
│   ├── skills_test.go
│   └── permissions_test.go
├── docs/
│   ├── TELEGRAM_GUIDE.md       # دليل Telegram
│   ├── BRIDGES_GUIDE.md        # دليل الجسور
│   ├── SKILLS_GUIDE.md         # دليل المهارات
│   └── ARCHITECTURE.md         # الهندسة المعمارية
├── go.mod
├── go.sum
├── .gitignore
├── LICENSE
├── CHANGELOG.md
└── CONTRIBUTING.md
```

## 🔧 المميزات الداخلية المتقدمة

### نظام الجسور (Bridges System)
- محاكاة تصرفات المستخدم الحقيقي
- إدارة الجلسات والكوكيز
- معالجة المصادقة التلقائية
- تجاوز تقنيات الكشف الحديثة
- دعم الوكلاء والشبكات الخاصة

### نظام المهارات (Skills System)
- تحميل المهارات الحية بدون إعادة بدء
- تدريب النظام على مهارات جديدة
- حفظ وتصدير المهارات
- إدارة تبعيات المهارات
- نظام الإصدارات للمهارات

### نظام الأذونات (Permissions System)
- قوائم بيضاء وسوداء
- تحكم حبيبي على العمليات
- تسجيل شامل للعمليات
- إشعارات تنبيهية للعمليات المشبوهة
- حدود معدل التشغيل

### نظام الذاكرة (Memory System)
- ذاكرة المشروع طويلة المدى
- ذاكرة الجلسة قصيرة المدى
- الفهرسة الذكية للمعلومات
- استرجاع سريع للمعلومات الهامة
- تحديث ديناميكي للذاكرة

## 📊 قياس الأداء والتحليل

```bash
# عرض إحصائيات الاستخدام
opencode-advanced --analytics

# تقرير الأداء
opencode-advanced --report

# تصدير البيانات
opencode-advanced --export-data /path/to/export
```

## 🧪 الاختبار

```bash
# تشغيل جميع الاختبارات
go test ./...

# اختبار محدد
go test ./internal/bridges/...

# مع التغطية
go test -cover ./...
```

## 🐛 استكشاف الأخطاء

```bash
# تشغيل بوضع التصحيح
opencode-advanced --debug

# حفظ السجلات
opencode-advanced --log-file debug.log

# وضع مفصل جداً
opencode-advanced --verbose
```

## 🔐 الأمان والخصوصية

- **تشفير البيانات**: جميع البيانات الحساسة مشفرة
- **إدارة السرية**: دعم متغيرات البيئة والملفات المشفرة
- **التدقيق الشامل**: تسجيل جميع العمليات الحساسة
- **عزل الجلسات**: كل جلسة معزولة عن الأخرى
- **حدود الموارد**: حدود على استهلاك المعالج والذاكرة

## 📚 التوثيق الإضافي

- [دليل Telegram الكامل](docs/TELEGRAM_GUIDE.md)
- [دليل الجسور والنماذج](docs/BRIDGES_GUIDE.md)
- [نظام المهارات المتقدم](docs/SKILLS_GUIDE.md)
- [الهندسة المعمارية التفصيلية](docs/ARCHITECTURE.md)
- [دليل تطوير البرامج الإضافية](docs/PLUGINS_GUIDE.md)

## 📦 إدارة الإصدارات

نتبع [Semantic Versioning](https://semver.org/lang/ar/)

## 🤝 المساهمة

نرحب بالمساهمات! يرجى اتباع [دليل المساهمة](CONTRIBUTING.md)

## 📄 الترخيص

MIT License - انظر [LICENSE](LICENSE) للتفاصيل

## 👨‍💻 المؤلف

تم تطويره بواسطة **xbitoi** 🚀

## 🙏 شكر خاص

- فريق OpenCode الأصلي للأساس القوي
- مجتمع Go العالمي
- المساهمون والمختبرون

---

**ملاحظة**: هذا المشروع قيد التطوير النشط. قد تحدث تغييرات كبيرة في الإصدارات القادمة.

**آخر تحديث**: 2026-04-25
