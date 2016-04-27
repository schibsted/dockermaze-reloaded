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
    name='heart-client',
    version='0.0.1',
    description='Heart solver client',
    author='SPT Infrastructure',
    author_email='tech-devops@schibsted.com',
    packages=find_packages(),
    entry_points={'console_scripts': [
        'heart-client = client.client:main'
    ]},
    include_package_data=True,
    install_requires=get_install_requirements('requirements.txt'),
    zip_safe=False,
)
