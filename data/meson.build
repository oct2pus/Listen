icon_sizes = ['16', '32', '48', '64', '128']

run_target('icons',
	command: 'icons/build.fish')

foreach i : icon_sizes
	install_data(
		join_paths('icons', i, meson.project_name() + '.png'),
		install_dir: join_paths(get_option('datadir'), 'icons', 'hicolor', i + 'x' + i, 'apps')
	)
endforeach

install_data(
    meson.project_name() + '.desktop',
    install_dir: join_paths(get_option('datadir'), 'applications')
)

install_data(
    meson.project_name() + '.appdata.xml',
    install_dir: join_paths(get_option('datadir'), 'metainfo')
)

install_data(
    meson.project_name() + '.contract',
    install_dir: join_paths(get_option('datadir'), 'contractor')
)