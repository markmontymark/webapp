## Playing with both custom go lang net/http and Go AppEngine

I'm toggling between Go App Engine dev and custom Go lang dev so need to set up environment a little differently between the two.

For GoAppEngine, see  ~/.bashrc file, but here it is:

	usegoappengine(){
		export GOAPPENGINEPATH=$HOME/vendor/go_appengine
		export PATH=$GOAPPENGINEPATH:$PATH
	}

For ad-hoc Go dev, source this directory's envrc file:

	source envrc

	cat envrc

	#!/bin/sh

	export LD_LIBRARY_PATH=$HOME/.gvm/pkgsets/go1.1.2/global/overlay/lib:
	export PATH=$HOME/.gvm/bin:$HOME/.gvm/pkgsets/go1.1.2/global/bin:$HOME/.gvm/gos/go1.1.2/bin:$HOME/.gvm/pkgsets/go1.1.2/global/overlay/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
	export DYLD_LIBRARY_PATH=$HOME/.gvm/pkgsets/go1.1.2/global/overlay/lib:
	export GVM_OVERLAY_PREFIX=$HOME/.gvm/pkgsets/go1.1.2/global/overlay
	export GOPATH=$HOME/.gvm/pkgsets/go1.1.2/global:$HOME/go-workspace
	export PKG_CONFIG_PATH=$HOME/.gvm/pkgsets/go1.1.2/global/overlay/lib/pkgconfig:
	export GVM_PATH_BACKUP=$HOME/.gvm/bin:$HOME/.gvm/pkgsets/go1.1.2/global/bin:$HOME/.gvm/gos/go1.1.2/bin:$HOME/.gvm/pkgsets/go1.1.2/global/overlay/bin:$HOME/.alternatives:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin


