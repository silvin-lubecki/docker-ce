# docker.fish - docker completions for fish shell
#
# This file is generated by gen_docker_fish_completions.py from:
# https://github.com/barnybug/docker-fish-completion
#
# To install the completions:
# mkdir -p ~/.config/fish/completions
# cp docker.fish ~/.config/fish/completions
#
# Completion supported:
# - parameters
# - commands
# - containers
# - images
# - repositories

function __fish_docker_no_subcommand --description 'Test if docker has yet to be given the subcommand'
    for i in (commandline -opc)
        if contains -- $i attach build commit cp diff events export history images import info insert inspect kill load login logs port ps pull push restart rm rmi run save search start stop tag top version wait
            return 1
        end
    end
    return 0
end

function __fish_print_docker_containers --description 'Print a list of docker containers' -a select
    switch $select
        case running
            docker ps -a --no-trunc | command awk 'NR>1' | command awk 'BEGIN {FS="  +"}; $5 ~ "^Up" {print $1 "\n" $(NF-1)}' | tr ',' '\n'
        case stopped
            docker ps -a --no-trunc | command awk 'NR>1' | command awk 'BEGIN {FS="  +"}; $5 ~ "^Exit" {print $1 "\n" $(NF-1)}' | tr ',' '\n'
        case all
            docker ps -a --no-trunc | command awk 'NR>1' | command awk 'BEGIN {FS="  +"}; {print $1 "\n" $(NF-1)}' | tr ',' '\n'
    end
end

function __fish_print_docker_images --description 'Print a list of docker images'
    docker images | command awk 'NR>1' | command grep -v '<none>' | command awk '{print $1":"$2}'
end

function __fish_print_docker_repositories --description 'Print a list of docker repositories'
    docker images | command awk 'NR>1' | command grep -v '<none>' | command awk '{print $1}' | command sort | command uniq
end

# common options
complete -c docker -f -n '__fish_docker_no_subcommand' -s D -l debug -d 'Enable debug mode'
complete -c docker -f -n '__fish_docker_no_subcommand' -s G -l group -d "Group to assign the unix socket specified by -H when running in daemon mode; use '' (the empty string) to disable setting of a group"
complete -c docker -f -n '__fish_docker_no_subcommand' -s H -l host -d 'tcp://host:port, unix://path/to/socket, fd://* or fd://socketfd to use in daemon mode. Multiple sockets can be specified'
complete -c docker -f -n '__fish_docker_no_subcommand' -l api-enable-cors -d 'Enable CORS headers in the remote API'
complete -c docker -f -n '__fish_docker_no_subcommand' -s b -l bridge -d "Attach containers to a pre-existing network bridge; use 'none' to disable container networking"
complete -c docker -f -n '__fish_docker_no_subcommand' -l bip -d "Use this CIDR notation address for the network bridge's IP, not compatible with -b"
complete -c docker -f -n '__fish_docker_no_subcommand' -s d -l daemon -d 'Enable daemon mode'
complete -c docker -f -n '__fish_docker_no_subcommand' -l dns -d 'Force docker to use specific DNS servers'
complete -c docker -f -n '__fish_docker_no_subcommand' -s e -l exec-driver -d 'Force the docker runtime to use a specific exec driver'
complete -c docker -f -n '__fish_docker_no_subcommand' -s g -l graph -d 'Path to use as the root of the docker runtime'
complete -c docker -f -n '__fish_docker_no_subcommand' -l icc -d 'Enable inter-container communication'
complete -c docker -f -n '__fish_docker_no_subcommand' -l ip -d 'Default IP address to use when binding container ports'
complete -c docker -f -n '__fish_docker_no_subcommand' -l ip-forward -d 'Disable enabling of net.ipv4.ip_forward'
complete -c docker -f -n '__fish_docker_no_subcommand' -l iptables -d "Disable docker's addition of iptables rules"
complete -c docker -f -n '__fish_docker_no_subcommand' -l mtu -d 'Set the containers network MTU; if no value is provided: default to the default route MTU or 1500 if no default route is available'
complete -c docker -f -n '__fish_docker_no_subcommand' -s p -l pidfile -d 'Path to use for daemon PID file'
complete -c docker -f -n '__fish_docker_no_subcommand' -s r -l restart -d 'Restart previously running containers'
complete -c docker -f -n '__fish_docker_no_subcommand' -s s -l storage-driver -d 'Force the docker runtime to use a specific storage driver'
complete -c docker -f -n '__fish_docker_no_subcommand' -s v -l version -d 'Print version information and quit'

# subcommands
# attach
complete -c docker -f -n '__fish_docker_no_subcommand' -a attach -d 'Attach to a running container'
complete -c docker -A -f -n '__fish_seen_subcommand_from attach' -l no-stdin -d 'Do not attach stdin'
complete -c docker -A -f -n '__fish_seen_subcommand_from attach' -l sig-proxy -d 'Proxify all received signal to the process (even in non-tty mode)'
complete -c docker -A -f -n '__fish_seen_subcommand_from attach' -a '(__fish_print_docker_containers running)' -d "Container"

# build
complete -c docker -f -n '__fish_docker_no_subcommand' -a build -d 'Build a container from a Dockerfile'
complete -c docker -A -f -n '__fish_seen_subcommand_from build' -l no-cache -d 'Do not use cache when building the image'
complete -c docker -A -f -n '__fish_seen_subcommand_from build' -s q -l quiet -d 'Suppress the verbose output generated by the containers'
complete -c docker -A -f -n '__fish_seen_subcommand_from build' -l rm -d 'Remove intermediate containers after a successful build'
complete -c docker -A -f -n '__fish_seen_subcommand_from build' -s t -l tag -d 'Repository name (and optionally a tag) to be applied to the resulting image in case of success'

# commit
complete -c docker -f -n '__fish_docker_no_subcommand' -a commit -d "Create a new image from a container's changes"
complete -c docker -A -f -n '__fish_seen_subcommand_from commit' -s a -l author -d 'Author (eg. "John Hannibal Smith <hannibal@a-team.com>"'
complete -c docker -A -f -n '__fish_seen_subcommand_from commit' -s m -l message -d 'Commit message'
complete -c docker -A -f -n '__fish_seen_subcommand_from commit' -l run -d 'Config automatically applied when the image is run. (ex: -run=\'{"Cmd": ["cat", "/world"], "PortSpecs": ["22"]}\')'
complete -c docker -A -f -n '__fish_seen_subcommand_from commit' -a '(__fish_print_docker_containers all)' -d "Container"

# cp
complete -c docker -f -n '__fish_docker_no_subcommand' -a cp -d 'Copy files/folders from the containers filesystem to the host path'

# diff
complete -c docker -f -n '__fish_docker_no_subcommand' -a diff -d "Inspect changes on a container's filesystem"
complete -c docker -A -f -n '__fish_seen_subcommand_from diff' -a '(__fish_print_docker_containers all)' -d "Container"

# events
complete -c docker -f -n '__fish_docker_no_subcommand' -a events -d 'Get real time events from the server'
complete -c docker -A -f -n '__fish_seen_subcommand_from events' -l since -d 'Show previously created events and then stream.'

# export
complete -c docker -f -n '__fish_docker_no_subcommand' -a export -d 'Stream the contents of a container as a tar archive'
complete -c docker -A -f -n '__fish_seen_subcommand_from export' -a '(__fish_print_docker_containers all)' -d "Container"

# history
complete -c docker -f -n '__fish_docker_no_subcommand' -a history -d 'Show the history of an image'
complete -c docker -A -f -n '__fish_seen_subcommand_from history' -l no-trunc -d "Don't truncate output"
complete -c docker -A -f -n '__fish_seen_subcommand_from history' -s q -l quiet -d 'Only show numeric IDs'
complete -c docker -A -f -n '__fish_seen_subcommand_from history' -a '(__fish_print_docker_images)' -d "Image"

# images
complete -c docker -f -n '__fish_docker_no_subcommand' -a images -d 'List images'
complete -c docker -A -f -n '__fish_seen_subcommand_from images' -s a -l all -d 'Show all images (by default filter out the intermediate image layers)'
complete -c docker -A -f -n '__fish_seen_subcommand_from images' -l no-trunc -d "Don't truncate output"
complete -c docker -A -f -n '__fish_seen_subcommand_from images' -s q -l quiet -d 'Only show numeric IDs'
complete -c docker -A -f -n '__fish_seen_subcommand_from images' -s t -l tree -d 'Output graph in tree format'
complete -c docker -A -f -n '__fish_seen_subcommand_from images' -s v -l viz -d 'Output graph in graphviz format'
complete -c docker -A -f -n '__fish_seen_subcommand_from images' -a '(__fish_print_docker_repositories)' -d "Repository"

# import
complete -c docker -f -n '__fish_docker_no_subcommand' -a import -d 'Create a new filesystem image from the contents of a tarball'

# info
complete -c docker -f -n '__fish_docker_no_subcommand' -a info -d 'Display system-wide information'

# insert
complete -c docker -f -n '__fish_docker_no_subcommand' -a insert -d 'Insert a file in an image'
complete -c docker -A -f -n '__fish_seen_subcommand_from insert' -a '(__fish_print_docker_images)' -d "Image"

# inspect
complete -c docker -f -n '__fish_docker_no_subcommand' -a inspect -d 'Return low-level information on a container'
complete -c docker -A -f -n '__fish_seen_subcommand_from inspect' -s f -l format -d 'Format the output using the given go template.'
complete -c docker -A -f -n '__fish_seen_subcommand_from inspect' -a '(__fish_print_docker_images)' -d "Image"
complete -c docker -A -f -n '__fish_seen_subcommand_from inspect' -a '(__fish_print_docker_containers all)' -d "Container"

# kill
complete -c docker -f -n '__fish_docker_no_subcommand' -a kill -d 'Kill a running container'
complete -c docker -A -f -n '__fish_seen_subcommand_from kill' -s s -l signal -d 'Signal to send to the container'
complete -c docker -A -f -n '__fish_seen_subcommand_from kill' -a '(__fish_print_docker_containers running)' -d "Container"

# load
complete -c docker -f -n '__fish_docker_no_subcommand' -a load -d 'Load an image from a tar archive'

# login
complete -c docker -f -n '__fish_docker_no_subcommand' -a login -d 'Register or Login to the docker registry server'
complete -c docker -A -f -n '__fish_seen_subcommand_from login' -s e -l email -d 'Email'
complete -c docker -A -f -n '__fish_seen_subcommand_from login' -s p -l password -d 'Password'
complete -c docker -A -f -n '__fish_seen_subcommand_from login' -s u -l username -d 'Username'

# logs
complete -c docker -f -n '__fish_docker_no_subcommand' -a logs -d 'Fetch the logs of a container'
complete -c docker -A -f -n '__fish_seen_subcommand_from logs' -s f -l follow -d 'Follow log output'
complete -c docker -A -f -n '__fish_seen_subcommand_from logs' -a '(__fish_print_docker_containers running)' -d "Container"

# port
complete -c docker -f -n '__fish_docker_no_subcommand' -a port -d 'Lookup the public-facing port which is NAT-ed to PRIVATE_PORT'
complete -c docker -A -f -n '__fish_seen_subcommand_from port' -a '(__fish_print_docker_containers running)' -d "Container"

# ps
complete -c docker -f -n '__fish_docker_no_subcommand' -a ps -d 'List containers'
complete -c docker -A -f -n '__fish_seen_subcommand_from ps' -s a -l all -d 'Show all containers. Only running containers are shown by default.'
complete -c docker -A -f -n '__fish_seen_subcommand_from ps' -l before -d 'Show only container created before Id or Name, include non-running ones.'
complete -c docker -A -f -n '__fish_seen_subcommand_from ps' -s l -l latest -d 'Show only the latest created container, include non-running ones.'
complete -c docker -A -f -n '__fish_seen_subcommand_from ps' -s n -d 'Show n last created containers, include non-running ones.'
complete -c docker -A -f -n '__fish_seen_subcommand_from ps' -l no-trunc -d "Don't truncate output"
complete -c docker -A -f -n '__fish_seen_subcommand_from ps' -s q -l quiet -d 'Only display numeric IDs'
complete -c docker -A -f -n '__fish_seen_subcommand_from ps' -s s -l size -d 'Display sizes'
complete -c docker -A -f -n '__fish_seen_subcommand_from ps' -l since -d 'Show only containers created since Id or Name, include non-running ones.'

# pull
complete -c docker -f -n '__fish_docker_no_subcommand' -a pull -d 'Pull an image or a repository from the docker registry server'
complete -c docker -A -f -n '__fish_seen_subcommand_from pull' -s t -l tag -d 'Download tagged image in repository'
complete -c docker -A -f -n '__fish_seen_subcommand_from pull' -a '(__fish_print_docker_images)' -d "Image"
complete -c docker -A -f -n '__fish_seen_subcommand_from pull' -a '(__fish_print_docker_repositories)' -d "Repository"

# push
complete -c docker -f -n '__fish_docker_no_subcommand' -a push -d 'Push an image or a repository to the docker registry server'
complete -c docker -A -f -n '__fish_seen_subcommand_from push' -a '(__fish_print_docker_images)' -d "Image"
complete -c docker -A -f -n '__fish_seen_subcommand_from push' -a '(__fish_print_docker_repositories)' -d "Repository"

# restart
complete -c docker -f -n '__fish_docker_no_subcommand' -a restart -d 'Restart a running container'
complete -c docker -A -f -n '__fish_seen_subcommand_from restart' -s t -l time -d 'Number of seconds to try to stop for before killing the container. Once killed it will then be restarted. Default=10'
complete -c docker -A -f -n '__fish_seen_subcommand_from restart' -a '(__fish_print_docker_containers running)' -d "Container"

# rm
complete -c docker -f -n '__fish_docker_no_subcommand' -a rm -d 'Remove one or more containers'
complete -c docker -A -f -n '__fish_seen_subcommand_from rm' -s f -l force -d 'Force removal of running container'
complete -c docker -A -f -n '__fish_seen_subcommand_from rm' -s l -l link -d 'Remove the specified link and not the underlying container'
complete -c docker -A -f -n '__fish_seen_subcommand_from rm' -s v -l volumes -d 'Remove the volumes associated to the container'
complete -c docker -A -f -n '__fish_seen_subcommand_from rm' -a '(__fish_print_docker_containers stopped)' -d "Container"

# rmi
complete -c docker -f -n '__fish_docker_no_subcommand' -a rmi -d 'Remove one or more images'
complete -c docker -A -f -n '__fish_seen_subcommand_from rmi' -s f -l force -d 'Force'
complete -c docker -A -f -n '__fish_seen_subcommand_from rmi' -a '(__fish_print_docker_images)' -d "Image"

# run
complete -c docker -f -n '__fish_docker_no_subcommand' -a run -d 'Run a command in a new container'
complete -c docker -A -f -n '__fish_seen_subcommand_from run' -s P -l publish-all -d 'Publish all exposed ports to the host interfaces'
complete -c docker -A -f -n '__fish_seen_subcommand_from run' -s a -l attach -d 'Attach to stdin, stdout or stderr.'
complete -c docker -A -f -n '__fish_seen_subcommand_from run' -s c -l cpu-shares -d 'CPU shares (relative weight)'
complete -c docker -A -f -n '__fish_seen_subcommand_from run' -l cidfile -d 'Write the container ID to the file'
complete -c docker -A -f -n '__fish_seen_subcommand_from run' -s d -l detach -d 'Detached mode: Run container in the background, print new container id'
complete -c docker -A -f -n '__fish_seen_subcommand_from run' -l dns -d 'Set custom dns servers'
complete -c docker -A -f -n '__fish_seen_subcommand_from run' -s e -l env -d 'Set environment variables'
complete -c docker -A -f -n '__fish_seen_subcommand_from run' -l entrypoint -d 'Overwrite the default entrypoint of the image'
complete -c docker -A -f -n '__fish_seen_subcommand_from run' -l expose -d 'Expose a port from the container without publishing it to your host'
complete -c docker -A -f -n '__fish_seen_subcommand_from run' -s h -l hostname -d 'Container host name'
complete -c docker -A -f -n '__fish_seen_subcommand_from run' -s i -l interactive -d 'Keep stdin open even if not attached'
complete -c docker -A -f -n '__fish_seen_subcommand_from run' -l link -d 'Add link to another container (name:alias)'
complete -c docker -A -f -n '__fish_seen_subcommand_from run' -l lxc-conf -d 'Add custom lxc options -lxc-conf="lxc.cgroup.cpuset.cpus = 0,1"'
complete -c docker -A -f -n '__fish_seen_subcommand_from run' -s m -l memory -d 'Memory limit (format: <number><optional unit>, where unit = b, k, m or g)'
complete -c docker -A -f -n '__fish_seen_subcommand_from run' -s n -l networking -d 'Enable networking for this container'
complete -c docker -A -f -n '__fish_seen_subcommand_from run' -l name -d 'Assign a name to the container'
complete -c docker -A -f -n '__fish_seen_subcommand_from run' -s p -l publish -d "Publish a container's port to the host (format: ip:hostPort:containerPort | ip::containerPort | hostPort:containerPort) (use 'docker port' to see the actual mapping)"
complete -c docker -A -f -n '__fish_seen_subcommand_from run' -l privileged -d 'Give extended privileges to this container'
complete -c docker -A -f -n '__fish_seen_subcommand_from run' -l rm -d 'Automatically remove the container when it exits (incompatible with -d)'
complete -c docker -A -f -n '__fish_seen_subcommand_from run' -l sig-proxy -d 'Proxify all received signal to the process (even in non-tty mode)'
complete -c docker -A -f -n '__fish_seen_subcommand_from run' -s t -l tty -d 'Allocate a pseudo-tty'
complete -c docker -A -f -n '__fish_seen_subcommand_from run' -s u -l user -d 'Username or UID'
complete -c docker -A -f -n '__fish_seen_subcommand_from run' -s v -l volume -d 'Bind mount a volume (e.g. from the host: -v /host:/container, from docker: -v /container)'
complete -c docker -A -f -n '__fish_seen_subcommand_from run' -l volumes-from -d 'Mount volumes from the specified container(s)'
complete -c docker -A -f -n '__fish_seen_subcommand_from run' -s w -l workdir -d 'Working directory inside the container'
complete -c docker -A -f -n '__fish_seen_subcommand_from run' -a '(__fish_print_docker_images)' -d "Image"

# save
complete -c docker -f -n '__fish_docker_no_subcommand' -a save -d 'Save an image to a tar archive'
complete -c docker -A -f -n '__fish_seen_subcommand_from save' -a '(__fish_print_docker_images)' -d "Image"

# search
complete -c docker -f -n '__fish_docker_no_subcommand' -a search -d 'Search for an image in the docker index'
complete -c docker -A -f -n '__fish_seen_subcommand_from search' -l no-trunc -d "Don't truncate output"
complete -c docker -A -f -n '__fish_seen_subcommand_from search' -s s -l stars -d 'Only displays with at least xxx stars'
complete -c docker -A -f -n '__fish_seen_subcommand_from search' -s t -l trusted -d 'Only show trusted builds'

# start
complete -c docker -f -n '__fish_docker_no_subcommand' -a start -d 'Start a stopped container'
complete -c docker -A -f -n '__fish_seen_subcommand_from start' -s a -l attach -d "Attach container's stdout/stderr and forward all signals to the process"
complete -c docker -A -f -n '__fish_seen_subcommand_from start' -s i -l interactive -d "Attach container's stdin"
complete -c docker -A -f -n '__fish_seen_subcommand_from start' -a '(__fish_print_docker_containers stopped)' -d "Container"

# stop
complete -c docker -f -n '__fish_docker_no_subcommand' -a stop -d 'Stop a running container'
complete -c docker -A -f -n '__fish_seen_subcommand_from stop' -s t -l time -d 'Number of seconds to wait for the container to stop before killing it.'
complete -c docker -A -f -n '__fish_seen_subcommand_from stop' -a '(__fish_print_docker_containers running)' -d "Container"

# tag
complete -c docker -f -n '__fish_docker_no_subcommand' -a tag -d 'Tag an image into a repository'
complete -c docker -A -f -n '__fish_seen_subcommand_from tag' -s f -l force -d 'Force'
complete -c docker -A -f -n '__fish_seen_subcommand_from tag' -a '(__fish_print_docker_images)' -d "Image"

# top
complete -c docker -f -n '__fish_docker_no_subcommand' -a top -d 'Lookup the running processes of a container'
complete -c docker -A -f -n '__fish_seen_subcommand_from top' -a '(__fish_print_docker_containers running)' -d "Container"

# version
complete -c docker -f -n '__fish_docker_no_subcommand' -a version -d 'Show the docker version information'

# wait
complete -c docker -f -n '__fish_docker_no_subcommand' -a wait -d 'Block until a container stops, then print its exit code'
complete -c docker -A -f -n '__fish_seen_subcommand_from wait' -a '(__fish_print_docker_containers running)' -d "Container"


