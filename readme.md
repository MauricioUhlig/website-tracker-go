# Website change tracker
Aplicação criada para rastrear e notificar por telegram quando o site for alterado

## Contexto 
Estive participando do processo seletivo para Mestrado em Computação Aplicada e estava ansioso para visualizar o resultado assim que postado. 
Logo tive a ideia de criar uma aplicação em Go para consultar a página do edital, procurando por um termo em específico. Se este temo não estivesse presente, significava que alguma atualização havia ocorrido.

Para realizar a notificação, optei por um bot de telegram que já possuo e utilizo a algum tempo para me notificar.

## Utilização
Para utilizar este scripts, basta alterar o arquivo `configs.yaml`, preenchendo com sua _API Key_ e _ChatId_. Também será necessário colocar a _URL_ que será rastreada e o termo que você procura (que deixará de existir quando atualizado).
