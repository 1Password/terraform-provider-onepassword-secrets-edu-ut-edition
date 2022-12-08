from python_terraform import *
import platform

# Create item to read
t = Terraform()
return_code, stdout, stderr, = t.apply(
    input=False, skip_plan=True, no_color=IsFlagged)
os.system('op item get newtitle2 --fields label=password > original_password.txt')
with open("original_password.txt", "r") as file:
    original_password = file.read()[:-1]

original_lines = []
with open("main.tf", "r") as file:
    for line in file.readlines():
        original_lines.append(line)

modified_lines = [x for x in original_lines]
modified_lines[16] = "    length        = 10\n"
with open("main.tf", "w") as file:
    file.writelines(modified_lines)

return_code, stdout, stderr, = t.apply(
    input=False, skip_plan=True, no_color=IsFlagged)
os.system('op item get newtitle2 --fields label=password > new_password.txt')
with open("new_password.txt", "r") as file:
    new_password = file.read()[:-1]

assert (original_password != new_password)
assert (len(new_password) == 10)

with open("main.tf", "w") as file:
    file.writelines(original_lines)

os.system('op item delete newtitle2')

if 'Windows' in platform.platform():
    os.system('del original_password.txt')
    os.system('del new_password.txt')
else:
    os.system('rm -f original_password.txt')
    os.system('rm -f new_password.txt')

print("pass")
