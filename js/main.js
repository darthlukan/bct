/* main.js
 *
 * @author Brian Tomlinson
 * @contact brian.tomlinson@linux.com
 *
 */
(function () {
    var React = require('react');
    var Card = React.createFactory(require('react-user-card'));

    var cardStl = {
        backgroundColor: 'rgba(40,40,40,0.8)',
        width: 'inherit',
        height: 'inherit',
        margin: 'inherit',
        padding: '1%'
    }

    var headerStl = {
        backgroundColor: '#901111',
        padding: '1%',
        color: '#ECEFF1',
        position: 'relative'
    }

    var contentStl = {
        color: 'inheriit',
        padding: '2%'
    }

    React.render(
            Card({
                userimg: '/img/me.jpg',
                username: 'DarthLukan',
                content: '"We\'re all humans, act accordingly."',
                cardStyle: cardStl,
                headerStyle: headerStl,
                contentStyle: contentStl
            }), document.getElementById('card')
    );
});
