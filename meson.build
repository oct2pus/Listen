project('moe.jade.oct2pus.listen')

run_command('meson/icons.fish')

run_target('build',
	command: 'meson/build.fish')

subdir('data')

comp = find_program('go')

env = environment()
env.set('GO111MODULE', 'on')

infile = 'src/main.go'
outfile = ''


mytarget = custom_target('moe.jade.oct2pus.listen',
  output : 'moe.jade.oct2pus.listen',
  input : 'src/main.go',
  command : [find_program('env'), 'GO111MODULE=on',comp,'build','-o', meson.project_name(), '@INPUT@',],
  install : true,
  install_dir : 'bin')