from python_terraform import *

t = Terraform()
# Create item and check return code
return_code, stdout, stderr, = t.apply(input=False, skip_plan=True, no_color=IsFlagged)
assert(return_code == 0)
# Delete item and check return code
return_code, stdout, stderr,  = t.apply(destroy=True, skip_plan=True)
assert(return_code == 0)

print("pass")
