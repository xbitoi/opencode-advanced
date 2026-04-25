# دليل المساهمة 🤝

شكراً لاهتمامك بالمساهمة في OpenCode Advanced!

## معايير الكود

### 1. الأسلوب والتنسيق
- استخدم `gofmt` لتنسيق الكود
- التزم بـ [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- استخدم أسماء وصفية للمتغيرات والدوال
- اكتب تعليقات واضحة للدوال والحزم

### 2. الاختبارات
- اكتب اختبارات شاملة لكل ميزة جديدة
- تأكد من تغطية المسارات الرئيسية
- اختبر الحالات الحدية والاستثناءات
- استخدم `go test -cover` للتحقق من التغطية

### 3. التوثيق
- أضف تعليقات توثيقية لجميع الحزم والدوال العامة
- حدّث README.md إذا أضفت مميزات جديدة
- وثّق التغييرات في CHANGELOG.md
- أضف أمثلة الاستخدام للميزات المعقدة

### 4. اختبار الجودة
```bash
# تنسيق الكود
gofmt -s -w ./...

# فحص الأخطاء
go vet ./...

# التحليل الثابت
golangci-lint run

# الاختبارات
go test -v -race -coverprofile=coverage.out ./...
```

## عملية المساهمة

### 1. Fork المستودع
```bash
gh repo fork xbitoi/opencode-advanced --clone
cd opencode-advanced
```

### 2. إنشاء فرع للميزة
```bash
git checkout -b feature/your-amazing-feature
```

### 3. تطوير الميزة
- اكتب الكود بحسب معايير المشروع
- أضف الاختبارات المناسبة
- وثّق التغييرات

### 4. الاختبار المحلي
```bash
# تشغيل الاختبارات
go test -v ./...

# اختبار البناء
go build -o opencode-advanced ./cmd/main.go

# الاختبار اليدوي
./opencode-advanced --help
```

### 5. Commit التغييرات
```bash
gh git commit -m "feat: add amazing feature

Detailed description of the feature.

Closes #123"
```

### 6. Push والـ Pull Request
```bash
git push origin feature/your-amazing-feature
gh pr create --title "Add amazing feature" --body "Description"
```

## معايير Commit

اتبع [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

### أنواع الـ Commits:
- `feat`: ميزة جديدة
- `fix`: إصلاح خلل
- `docs`: تغييرات التوثيق
- `style`: تنسيق الكود
- `refactor`: إعادة تنظيم الكود
- `test`: إضافة اختبارات
- `chore`: تحديثات الأدوات

### أمثلة:
```
feat(telegram): add command scheduling
fix(bridges): resolve gemini connection timeout
docs(skills): update skill creation guide
test(permissions): add integration tests
```

## الإبلاغ عن الأخطاء

قبل الإبلاغ عن خطأ:
1. تأكد من أنه لم يتم الإبلاغ عنه مسبقاً
2. اختبر مع أحدث إصدار
3. اجمع معلومات التشخيص

### تنسيق الإبلاغ:
```markdown
## الوصف
وصف واضح للمشكلة

## خطوات إعادة الإنتاج
1. ...
2. ...
3. ...

## السلوك المتوقع
...

## السلوك الفعلي
...

## معلومات البيئة
- OS: Windows 11
- Go version: 1.24
- Branch: main

## السجلات والأخطاء
```

## طلبات الميزات

### قبل البدء:
1. ناقش الميزة في Issue أولاً
2. تأكد من توافقها مع رؤية المشروع
3. احصل على موافقة المشرفين

### تنسيق الطلب:
```markdown
## الوصف
وصف واضح للميزة المطلوبة

## الحالة الحالية
كيف يعمل الآن

## الحالة المرغوبة
كيف يجب أن يعمل

## الفوائد
- فائدة 1
- فائدة 2

## تصميم محتمل
أي أفكار حول التنفيذ
```

## مراجعة الأكواد (Code Review)

عند مراجعة PR:

### نقاط التركيز:
- ✅ يتبع معايير الكود
- ✅ اختبارات شاملة
- ✅ توثيق واضح
- ✅ لا تضاربات
- ✅ الأداء مقبول
- ✅ الأمان كافي

### الملاحظات البناءة:
- كن محترماً وموضوعياً
- اشرح السبب وراء كل مرة
- اقترح بدائل
- امدح الأشياء الجيدة

## قناة الاتصال

- **Issues**: لطلب الميزات والإبلاغ عن الأخطاء
- **Discussions**: للنقاشات العامة والأسئلة
- **Pull Requests**: لطلب الدمج

## الترخيص

بالمساهمة، توافق على ترخيص مشروعك تحت MIT License.

## الشكر

شكراً لمساهمتك في جعل OpenCode Advanced أفضل! 🙏
