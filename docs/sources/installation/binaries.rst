:title: Installation from Binaries
:description: This instruction set is meant for hackers who want to try out Docker on a variety of environments.
:keywords: binaries, installation, docker, documentation, linux

.. _binaries:

Binaries
========

.. include:: install_header.inc

**This instruction set is meant for hackers who want to try out Docker
on a variety of environments.**

Before following these directions, you should really check if a packaged version
of Docker is already available for your distribution.  We have packages for many
distributions, and more keep showing up all the time!

Check Your Kernel
-----------------

Your host's Linux kernel must meet the Docker :ref:`kernel`

Check for User Space Tools
--------------------------

You must have a working installation of the `lxc <http://linuxcontainers.org>`_ utilities and library.

Get the docker binary:
----------------------

.. code-block:: bash

    wget https://get.docker.io/builds/Linux/x86_64/docker-latest -O docker
    chmod +x docker


Run the docker daemon
---------------------

.. code-block:: bash

    # start the docker in daemon mode from the directory you unpacked
    sudo ./docker -d &

Upgrades
--------

To upgrade your manual installation of Docker, first kill the docker daemon:

.. code-block:: bash

   killall docker

Then follow the regular installation steps.


Run your first container!
-------------------------

.. code-block:: bash

    # check your docker version
    sudo ./docker version

    # run a container and open an interactive shell in the container
    sudo ./docker run -i -t ubuntu /bin/bash



Continue with the :ref:`hello_world` example.
