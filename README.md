# AnGGDev_GitLearning
## Chào mừng các cốt đến với khóa học lệnh git
### Clone the repo
dùng để tải xuống repo
```
git clone https://github.com/ducan172002/AnGGDev_GitLearning.git
```
### Init
Dùng để cài đặt git cho project
```
git init
```
### Add
Thêm tất cả các file cần up lên github
```
git add .
```
Thêm 1 file cần up lên github
```
git add Ten_file
```
### Branch
Nhánh của repo
Kiểm tra có những nhánh nào
```
git branch
```
Thêm hoặc truy cập vào 1 nhánh
```
git branch -M main
```
### Commit
Dùng để ghi chú ngắn gọn cho project
```
git commit -m "init commit"
```
### Remote
Kết nối Repository với project
```
git remote add origin https://github.com/ducan172002/AnGGDev_GitLearning.git
```
### Push
Đẩy project lên github
```
git push origin main
```
## Với 1 project tự mình up lên git
```
git init
```
```
git add .
```
```
git commit -m "init commit"
```
```
git remote add origin https://github.com/ducan172002/AnGGDev_GitLearning.git
```
```
git push origin main
```
## Đã có sẵn repo trên github
```
git remote add origin https://github.com/ducan172002/AnGGDev_GitLearning.git
```
```
git commit -m "init commit"
```
```
git push -u origin main
```
## Một vài lệnh khác
### Status
Kiểm tra các file có được thêm để up lên git chưa
Màu xanh: đã được thêm
Màu đỏ: chưa được thêm
```
git status
```
muốn hiện thị thông tin ngắn gọn hơn
```
git status -s
```
### Log
Xem lại lịch sử commit
```
git log
```
Lọc theo ngày
```
git log --after="2019-1-1" --before="2019-12-31"
```
Lọc theo người dùng
```
git log --oneline --author="XuanThuLab"
```
lọc theo thông tin ghi chú về commit
```
git log --oneline --grep="keyword"
```
Lọc các commit bình thường (tham số --no-merges) và các commit do gộp nhánh (tham số --merges)
```
git log --merges
```
