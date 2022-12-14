from python_terraform import *
import platform

# Create item to read
os.system('op item create --category="Login" --title="login" --vault="test" password=test_password > temp.txt')
t = Terraform()
# read item
return_code, stdout, stderr, = t.plan(input=False)
# find output from terraform output
index = stdout.find("login_secret")
# test equivalence to expeceted
assert (stdout[index + len("login_secret") + 4:index +
        len("login_secret") + 4 + 13] == "test_password")
# cleanup
os.system('op item delete login')

if 'Windows' in platform.platform():
    os.system('del temp.txt')
else:
    os.system('rm -f temp.txt')

print("pass")
