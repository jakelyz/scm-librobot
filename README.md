
# Table of Contents

1.  [scm-librobot](#orgb4a29b6)
    1.  [Creating the shared object](#org7e839c4)


<a id="orgb4a29b6"></a>

# scm-librobot

This repository contains a small go file that exports functions from the [robotgo](https://github.com/go-vgo/robotgo) project,
and a R7RS Scheme library to call those functions. This currently supports [ypsilon](https://github.com/fujita-y/ypsilon) scheme.


<a id="org7e839c4"></a>

## Creating the shared object

Ensure that you have go installed. The shared obj can be created via Make:

    # Darwin
    make build-darwin
    
    # M1, etc
    make ARCH=arm64 build-darwin
    
    # Linux
    make build-linux

