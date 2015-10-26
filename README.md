A Fast and Flexible CMS built with love by [weldpua2008](https://github.com/weldpua2008) and [friends](https://github.com/weldpua2008/hugo-cms/graphs/contributors) in [Go][].
It does edit a TOML, YAML and JSON metadata that uses in different static site generators (main target is a [Hugo]: https://gohugo.io/)

## Overview

HugoCMS is a CMS for static site written in [Go][].

#### Supported Architectures

Binaries will be provided for Windows, Linux, FreeBSD, NetBSD and OS&nbsp;X (Darwin) for x64, i386 and ARM architectures.

HugoCMS may also be compiled from source wherever the Go compiler tool chain can run, e.g. for other operating systems including DragonFly BSD, OpenBSD, Plan&nbsp;9 and Solaris.

## Contribute to HugoCMS

We welcome contributions to HugoCMS of any kind including documentation, themes, organization, tutorials, blog posts, bug reports, issues, feature requests, feature implementation, pull requests, answering questions on the forum, helping to manage issues, etc. 
The Hugo community and maintainers are very active and helpful and the project benefits greatly from this activity.

### Clone the HugoCMS Project (Contributor)

1. Make sure your local environment has the following software installed:

    * [Git](https://git-scm.com/)    
    * [Go][] 1.4+
    
2. [Fork the HugoCMS project on GitHub](https://github.com/weldpua2008/hugo-cms/).
3. Clone your fork:

        git clone https://github.com/YOURNAME/hugo-cms

4. Change into the `hugo-cms` directory:
        cd hugo-cms

5. Install the Hugo project’s package dependencies:

        go get -u -v github.com/weldpua2008/hugo-cms
6. Use a symbolic link to add your locally cloned Hugo repository to your `$GOPATH`, assuming you prefer doing development work outside of `$GOPATH`:

    ``` bash
    rm -rf "$GOPATH/src/github.com/weldpua2008/hugo-cms"
    ln -s `pwd` "$GOPATH/src/github.com/weldpua2008/hugo-cms"
    ```

    Go expects all of your libraries to be found in`$GOPATH`.


### Build and Install the Binaries from Source (Advanced Install)

Add HugoCMS and its package dependencies to your go `src` directory.

    go get -v github.com/weldpua2008/hugo-cms

Once the `get` completes, you should find your new `hugocms` (or `hugocms.exe`) executable sitting inside `$GOPATH/bin/`.

To update HugoCMS’s dependencies, use `go get` with the `-u` option.

    go get -u -v github.com/weldpua2008/hugo-cms

[Go]: https://golang.org/