from setuptools import setup, find_packages
import inspect
import os

__location__ = os.path.join(
    os.getcwd(), os.path.dirname(inspect.getfile(inspect.currentframe()))
)


def get_install_requirements(path):
    content = open(os.path.join(__location__, path)).read()
    requires = [req for req in content.split('\\n') if req != '']
    return requires

setup(
    name='maze-client',
    version='0.0.1',
    description='Maze solver client',
    author='Schibsted Products and Technology',
    author_email='big.ideas@schibsted.com',
    packages=find_packages(),
    entry_points={'console_scripts': [
        'maze-client = client.client:main'
    ]},
    include_package_data=True,
    install_requires=get_install_requirements('requirements.txt'),
    zip_safe=False,
)
