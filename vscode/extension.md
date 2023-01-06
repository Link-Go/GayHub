## Vscode自用拓展列表

获取拓展列表 - 需要先下载安装`vscode`
```bash
# unix
code --list-extensions | xargs -L 1 echo code --install-extension

# windows
code --list-extensions | % { "code --install-extension $_" }
```



### 拓展列表
```
code --install-extension aldijav.golangwithdidi
code --install-extension alefragnani.project-manager
code --install-extension bierner.markdown-preview-github-styles
code --install-extension eamodio.gitlens
code --install-extension formulahendry.code-runner
code --install-extension GitHub.remotehub
code --install-extension golang.go
code --install-extension metaseed.MetaJump
code --install-extension mhutchie.git-graph
code --install-extension ms-azuretools.vscode-docker
code --install-extension MS-CEINTL.vscode-language-pack-zh-hans
code --install-extension ms-python.isort
code --install-extension ms-python.python
code --install-extension ms-python.vscode-pylance
code --install-extension ms-toolsai.jupyter
code --install-extension ms-toolsai.jupyter-keymap
code --install-extension ms-toolsai.jupyter-renderers
code --install-extension ms-toolsai.vscode-jupyter-cell-tags
code --install-extension ms-toolsai.vscode-jupyter-slideshow
code --install-extension ms-vscode-remote.remote-ssh
code --install-extension ms-vscode-remote.remote-ssh-edit
code --install-extension ms-vscode.azure-repos
code --install-extension ms-vscode.remote-explorer
code --install-extension ms-vscode.remote-repositories
code --install-extension Natizyskunk.sftp
code --install-extension nhoizey.gremlins
code --install-extension premparihar.gotestexplorer
code --install-extension quicktype.quicktype
code --install-extension SirTori.indenticator
code --install-extension streetsidesoftware.code-spell-checker
code --install-extension stuart.unique-window-colors
code --install-extension techer.open-in-browser
code --install-extension unbug.codelf
code --install-extension zxh404.vscode-proto3
```


### 直接复制
```
# 电脑A 复制到电脑B 就相当于安装成功
Windows %USERPROFILE%\.vscode\extensions
Mac ~/.vscode/extensions
Linux ~/.vscode/extensions
```