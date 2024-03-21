import subprocess
import sys
import os

# This Python script fetches a file from a remote git repo
# Arguments are:
# 1. The URL of the git repo
# 2. The name of the branch to fetch
# 3. The name of the specification file that is used to whitelist files

# Check if the number of arguments is correct
if len(sys.argv) != 4:
    print("Illegal number of parameters")
    sys.exit(1)

# Assign arguments to variables
url = sys.argv[1]
branch = sys.argv[2]
spec_file = sys.argv[3]

try:
    # Initialize empty Git repository
    subprocess.run(['git', 'init'], check=True)
    
    # Enable sparse checkout
    subprocess.run(['git', 'config', 'core.sparseCheckout', 'true'], check=True)
    
    # Add whitelist to sparse-checkout
    with open('.git/info/sparse-checkout', 'a') as sparse_checkout:
        with open(spec_file, 'r') as whitelist:
            sparse_checkout.write(whitelist.read())
    
    # Add remote repository
    subprocess.run(['git', 'remote', 'add', 'origin', url], check=True)
    
    # Pull from the remote repository
    subprocess.run(['git', 'pull', 'origin', branch], check=True)
    
    # Remove the .git directory
    if os.name == 'nt':  # Windows
        subprocess.run(['rmdir', '/S', '/Q', '.git'], shell=True, check=True)
    else:
        subprocess.run(['rm', '-rf', '.git'], check=True)

except subprocess.CalledProcessError as e:
    print(f"An error occurred: {e}")
    sys.exit(1)